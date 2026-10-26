#!/usr/bin/env python3
"""Publish a release SDK archive as a non-runnable OCI artifact in GHCR."""
import argparse
import base64
import hashlib
import json
import os
import re
import urllib.error
import urllib.parse
import urllib.request

parser = argparse.ArgumentParser()
parser.add_argument('--version', required=True)
args = parser.parse_args()
version = args.version.removeprefix('v')
if not re.fullmatch(r'\d+\.\d+\.\d+', version):
    raise SystemExit('Versão inválida.')
tag = 'v' + version
repository = os.environ.get('GITHUB_REPOSITORY', 'GaabsDev/fiscal-tributario-brasil')
if repository.lower() != 'gaabsdev/fiscal-tributario-brasil':
    raise SystemExit('Repositório diferente do destino autorizado.')
token = os.environ.get('GITHUB_TOKEN') or os.environ.get('JAILSON_GITHUB_TOKEN')
if not token:
    raise SystemExit('Token ausente.')
actor = os.environ.get('GITHUB_ACTOR', 'GaabsDev')
registry_name = repository.lower()
artifact_type = 'application/vnd.gaabsdev.fiscal-tributario-brasil.sdk.v1'

def request(url, method='GET', data=None, headers=None):
    req = urllib.request.Request(url, data=data, method=method, headers={'User-Agent': 'Fiscal-Tributario-Brasil-OCI', **(headers or {})})
    return urllib.request.urlopen(req, timeout=60)

def api(path):
    with request('https://api.github.com' + path, headers={'Authorization': 'Bearer ' + token, 'Accept': 'application/vnd.github+json'}) as response:
        return json.load(response)

release = api('/repos/' + repository + '/releases/tags/' + tag)
filename = 'fiscal-tributario-brasil-' + tag + '.tar.gz'
asset = next(a for a in release['assets'] if a['name'] == filename)
with request(asset['browser_download_url']) as response:
    archive = response.read()
archive_hash = hashlib.sha256(archive).hexdigest()
if asset.get('digest') != 'sha256:' + archive_hash:
    raise SystemExit('Hash do arquivo não corresponde ao asset da release.')
ref = api('/repos/' + repository + '/git/ref/tags/' + tag)['object']
if ref['type'] == 'tag':
    ref = api('/repos/' + repository + '/git/tags/' + ref['sha'])['object']
revision = ref['sha']

basic = base64.b64encode((actor + ':' + token).encode()).decode()
query = urllib.parse.urlencode({'service': 'ghcr.io', 'scope': 'repository:' + registry_name + ':pull,push'})
with request('https://ghcr.io/token?' + query, headers={'Authorization': 'Basic ' + basic}) as response:
    bearer = json.load(response)['token']
headers = {'Authorization': 'Bearer ' + bearer}
root = 'https://ghcr.io/v2/' + registry_name

def blob(data):
    digest = 'sha256:' + hashlib.sha256(data).hexdigest()
    try:
        with request(root + '/blobs/' + digest, method='HEAD', headers=headers):
            return digest
    except urllib.error.HTTPError as error:
        if error.code != 404:
            raise
    with request(root + '/blobs/uploads/', method='POST', data=b'', headers=headers) as response:
        location = urllib.parse.urljoin(root + '/', response.headers['Location'])
    if urllib.parse.urlparse(location).hostname != 'ghcr.io':
        raise SystemExit('Destino de upload inesperado.')
    separator = '&' if '?' in location else '?'
    with request(location + separator + urllib.parse.urlencode({'digest': digest}), method='PUT', data=data, headers={**headers, 'Content-Type': 'application/octet-stream'}) as response:
        if response.status != 201:
            raise SystemExit('Upload incompleto.')
    return digest

manifest = {
    'schemaVersion': 2,
    'mediaType': 'application/vnd.oci.image.manifest.v1+json',
    'artifactType': artifact_type,
    'config': {'mediaType': 'application/vnd.oci.empty.v1+json', 'digest': 'sha256:' + hashlib.sha256(b'{}').hexdigest(), 'size': 2},
    'layers': [{'mediaType': 'application/gzip', 'digest': 'sha256:' + archive_hash, 'size': len(archive), 'annotations': {'org.opencontainers.image.title': filename}}],
    'annotations': {
        'org.opencontainers.image.title': 'fiscal-tributario-brasil',
        'org.opencontainers.image.version': version,
        'org.opencontainers.image.source': 'https://github.com/' + repository,
        'org.opencontainers.image.revision': revision,
        'org.opencontainers.image.licenses': 'MIT',
        'org.opencontainers.image.description': 'Fiscal Tributário Brasil: arquivo do SDK Go e documentação. Artefato OCI sem sistema operacional ou executável. Baixe com oras pull.',
    },
}
accept = {'Accept': manifest['mediaType']}
try:
    with request(root + '/manifests/' + version, headers={**headers, **accept}) as response:
        previous = json.load(response)
    if previous.get('artifactType') != artifact_type or previous.get('layers', [{}])[0].get('digest') != 'sha256:' + archive_hash:
        raise SystemExit('A versão já existe com conteúdo diferente; não será sobrescrita.')
    print('Artefato já publicado com o mesmo conteúdo.')
except urllib.error.HTTPError as error:
    if error.code != 404:
        raise
    blob(b'{}')
    blob(archive)
    body = json.dumps(manifest, sort_keys=True, separators=(',', ':')).encode()
    with request(root + '/manifests/' + version, method='PUT', data=body, headers={**headers, 'Content-Type': manifest['mediaType']}) as response:
        print('Manifesto publicado:', response.status, response.headers.get('Docker-Content-Digest'))
with request(root + '/manifests/' + version, headers={**headers, **accept}) as response:
    verified = json.load(response)
assert verified['artifactType'] == artifact_type
assert verified['config']['mediaType'] == 'application/vnd.oci.empty.v1+json'
assert verified['layers'][0]['digest'] == 'sha256:' + archive_hash
print(json.dumps({'package': 'ghcr.io/' + registry_name + ':' + version, 'artifact_type': artifact_type, 'archive_sha256': archive_hash, 'runnable_image': False}, ensure_ascii=False))

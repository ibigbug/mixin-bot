import os
import sys

from azure.storage.blob import BlockBlobService
from azure.common import AzureMissingResourceHttpError

account_name = os.environ.get('azure_storage_account_name')
account_key = os.environ.get('azure_storage_account_key')
container_name = os.environ.get('azure_storage_container_name')
block_blob_service = BlockBlobService(account_name, account_key)

target_folder = sys.argv[1]

for f in os.listdir(target_folder):
    if not f.startswith('.'):
        full_path = os.path.join(target_folder, f)
        try:
            block_blob_service.get_blob_metadata(
                container_name=container_name, blob_name=f)
        except AzureMissingResourceHttpError as ex:
            block_blob_service.create_blob_from_path(
                container_name, f, full_path)

for b in block_blob_service.list_blobs(container_name):
    print(b.name)

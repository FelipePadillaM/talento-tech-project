import boto3
from typing import List
from botocore.exceptions import BotoCoreError, ClientError
from fastapi import HTTPException

class S3Service:
    def __init__(self, bucket_name: str):
        self.s3 = boto3.client('s3')
        self.bucket_name = bucket_name

    def list_files(self) -> List[str]:
        try:
            response = self.s3.list_objects_v2(Bucket=self.bucket_name)
            if 'Contents' in response:
                return [obj['Key'] for obj in response['Contents']]
            return []
        except (BotoCoreError, ClientError) as e:
            raise HTTPException(status_code=500, detail=f"Error al conectar con S3: {str(e)}")

    def upload_fileobj(self, fileobj, filename: str):
        try:
            self.s3.upload_fileobj(fileobj, self.bucket_name, filename)
        except (BotoCoreError, ClientError) as e:
            raise HTTPException(status_code=500, detail=f"Error al subir archivo a S3: {str(e)}") 
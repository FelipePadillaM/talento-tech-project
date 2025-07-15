import boto3
from typing import List

class S3Service:
    def __init__(self, bucket_name: str):
        self.s3 = boto3.client('s3')
        self.bucket_name = bucket_name

    def list_files(self) -> List[str]:
        response = self.s3.list_objects_v2(Bucket=self.bucket_name)
        if 'Contents' in response:
            return [obj['Key'] for obj in response['Contents']]
        return [] 
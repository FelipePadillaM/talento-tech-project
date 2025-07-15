from fastapi import APIRouter, Depends
from app.services.s3_service import S3Service

router = APIRouter()

def get_s3_service():
    # Aquí podrías obtener el bucket de config/env
    return S3Service(bucket_name="nombre-de-tu-bucket")

@router.get("/files", summary="Listar archivos del bucket")
def list_files(s3_service: S3Service = Depends(get_s3_service)):
    return {"files": s3_service.list_files()} 
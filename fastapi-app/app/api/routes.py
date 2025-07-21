from fastapi import APIRouter, Depends, File, UploadFile
from fastapi.responses import JSONResponse
from app.services.s3_service import S3Service
from app.config import BUCKET_NAME

router = APIRouter()

def get_s3_service():
    return S3Service(bucket_name=BUCKET_NAME)

@router.get("/files", summary="Listar archivos del bucket")
def list_files(s3_service: S3Service = Depends(get_s3_service)):
    return {"files": s3_service.list_files()}

@router.get("/dummy-files", summary="Lista dummy de archivos del bucket")
def dummy_files():
    return {
        "files": [
            {"filename": "archivo1.txt", "size": 12345, "url": "https://dummy-bucket.s3.amazonaws.com/archivo1.txt"},
            {"filename": "foto.png", "size": 54321, "url": "https://dummy-bucket.s3.amazonaws.com/foto.png"},
            {"filename": "documento.pdf", "size": 9999, "url": "https://dummy-bucket.s3.amazonaws.com/documento.pdf"}
        ]
    }

@router.post("/upload-file", summary="Subir un archivo al bucket S3")
def upload_file(file: UploadFile = File(...), s3_service: S3Service = Depends(get_s3_service)):
    try:
        s3_service.upload_fileobj(file.file, file.filename)
        return {"message": f"Archivo '{file.filename}' subido exitosamente."}
    except Exception as e:
        return JSONResponse(status_code=500, content={"error": str(e)}) 
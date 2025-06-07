from fastapi import FastAPI, UploadFile, HTTPException
from fastapi.middleware.cors import CORSMiddleware
import librosa
import numpy as np
import tempfile
import os

app = FastAPI()

# Allow all origins for frontend access
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_methods=["POST"],
    allow_headers=["*"],
)

def detect_pitch(audio_data: np.ndarray, sr: int = 22050) -> list[str]:
    """Detect musical notes from audio using YIN algorithm."""
    pitches = librosa.yin(
        audio_data,
        fmin=librosa.note_to_hz('C2'),
        fmax=librosa.note_to_hz('C7'),
        sr=sr
    )
    return [librosa.hz_to_note(p) for p in pitches if not np.isnan(p)]

@app.post("/detect")
async def process_audio(file: UploadFile):
    if not file.content_type.startswith("audio/"):
        raise HTTPException(400, "Uploaded file must be audio (WAV/MP3)")

    try:
        # Save uploaded file temporarily
        with tempfile.NamedTemporaryFile(delete=False, suffix=".wav") as tmp:
            tmp.write(await file.read())
            tmp_path = tmp.name

        # Load audio and detect pitches
        y, sr = librosa.load(tmp_path, sr=None)
        os.unlink(tmp_path)  # Clean up

        pitches = detect_pitch(y, sr)
        return {"pitches": pitches}

    except Exception as e:
        raise HTTPException(500, f"Audio processing failed: {str(e)}")

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=5000)

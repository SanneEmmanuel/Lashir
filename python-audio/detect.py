# python-audio/detect.py
import librosa
import numpy as np
import sys

def detect_pitch(audio_path: str) -> list:
    y, sr = librosa.load(audio_path)
    pitches = librosa.yin(y, fmin=librosa.note_to_hz('C2'), fmax=librosa.note_to_hz('C7'))
    return [librosa.hz_to_note(p) for p in pitches if not np.isnan(p)]

if __name__ == "__main__":
    print(detect_pitch(sys.argv[1]))

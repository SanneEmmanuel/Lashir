import numpy as np  
from scipy.signal import find_peaks  

def detect_pitch(audio: np.ndarray, sr: int) -> list[str]:  
    """E pitch extraction via harmonic peak analysis"""  
    spectrum = np.fft.fft(audio)  
    peaks, _ = find_peaks(np.abs(spectrum), prominence=0.5)  
    return [hz_to_note(peak * sr / len(audio)) for peak in peaks]  

from music21 import *  

def generate_notation(pitches: list, key: str = "C") -> str:  
    """Render sheet music with solfa glyphs"""  
    s = stream.Stream()  
    for p in pitches:  
        n = note.Note(p)  
        n.addLyric(solfa_mapping(p, key))  # 🎵 Do, Re, Mi...  
    return s.write('lilypond')  

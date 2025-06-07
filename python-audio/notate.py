# python-music21/notate.py
from music21 import *
import sys

def generate_sheet_music(pitches: list) -> str:
    s = stream.Stream()
    for p in pitches:
        n = note.Note(p)
        n.lyric = convert_to_solfa(p)  # Implement your solfa mapping
        s.append(n)
    s.write('pdf', 'output.pdf')
    return 'output.pdf'

if __name__ == "__main__":
    pitches = eval(sys.argv[1])  # Passed from Go
    print(generate_sheet_music(pitches))

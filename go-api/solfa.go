// go-api/solfa.go
package main

func convertToSolfa(note string) string {
    solfa := map[string]string{
        "C": "Do", "D": "Re", "E": "Mi",
        "F": "Fa", "G": "Sol", "A": "La", "B": "Ti",
    }
    if len(note) > 0 {
        return solfa[note[:1]] // Extract base note (e.g., "C" from "C4")
    }
    return ""
}

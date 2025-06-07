# Lashir API

## 🎼 Real-Time Audio to Sheet Music Conversion

**Live Demo:** [https://lashir-api.onrender.com](https://lashir-api.onrender.com)

Lashir is a powerful API that transforms audio into beautifully notated sheet music with solfege (Do, Re, Mi) annotations. Perfect for musicians, educators, and developers.

## ✨ Key Features

- **Instant Conversion**: Upload audio files or stream live
- **Professional Notation**: Clean, readable sheet music output
- **Solfa Support**: Automatic Do-Re-Mi labeling
- **RESTful API**: Easy integration with any application

## 🚀 Quick Start

### API Endpoints

**1. File Upload Conversion**
```
POST https://lashir-api.onrender.com/upload
```
- **Content-Type**: `multipart/form-data`
- **Body**: `{ "audio": file }` (WAV/MP3)

**Example (cURL):**
```bash
curl -X POST -F "audio=@recording.wav" https://lashir-api.onrender.com/upload
```

**Response:**
```json
{
  "pdf_url": "https://lashir-api.onrender.com/output.pdf",
  "solfa_notation": ["Do", "Re", "Mi"]
}
```

## 🛠 Integration Guide

### Web Application
```javascript
async function convertAudio(file) {
  const formData = new FormData();
  formData.append('audio', file);
  
  const response = await fetch('https://lashir-api.onrender.com/upload', {
    method: 'POST',
    body: formData
  });
  
  return await response.json();
}
```

### Mobile Apps
```dart
// Flutter example
Future<Map> convertAudio(File file) async {
  var request = http.MultipartRequest(
    'POST',
    Uri.parse('https://lashir-api.onrender.com/upload')
  );
  request.files.add(await http.MultipartFile.fromPath('audio', file.path));
  var response = await request.send();
  return jsonDecode(await response.stream.bytesToString());
}
```

## 📊 Technical Specifications

- **Max File Size**: 10MB (WAV/MP3)
- **Processing Time**: ~5-15 seconds (depending on length)
- **Output Formats**: PDF, JSON (with raw pitch data)

## 💡 Example Use Cases

1. **Music Education Apps**: Auto-generate exercises from student recordings
2. **Practice Tools**: Convert improvisations to readable notation
3. **Composition Aids**: Quickly transcribe musical ideas

## ⚠️ Current Limitations

- Works best with monophonic (single-note) audio
- Optimal results require clear recordings
- Free tier has rate limits (upgrade for production)

## 📜 License

Open-source under the Mystic Public License (MPL)

---

**Happy transcribing!** For support, contact wa.me/+2348109995000

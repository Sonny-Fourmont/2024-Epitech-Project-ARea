import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

class NewAppletScreen extends StatefulWidget {
  const NewAppletScreen({super.key});

  @override
  _NewAppletScreenState createState() => _NewAppletScreenState();
}

class _NewAppletScreenState extends State<NewAppletScreen> {
  final Dio dio = Dio();

  final TextEditingController ifController = TextEditingController();
  final TextEditingController thatController = TextEditingController();
  final TextEditingController ifTypeController = TextEditingController();
  final TextEditingController thatTypeController = TextEditingController();

  void postData() async {
    final apiUrl = dotenv.env['API_APPLET_URL'];
    if (apiUrl == null) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('L\'URL de l\'API n\'est pas configurée')),
      );
      return;
    }
    try {
      Response response = await dio.post(
        apiUrl,
        data: {
          "if": ifController.text,
          "ifType": ifTypeController.text,
          "isOn": true,
          "that": thatController.text,
          "thatType": thatTypeController.text,
        },
      );
      print('if text ${ifController.text}');
      print('Response data: ${response.data}');
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Applet créé avec succès !')),
      );
    } catch (e) {
      print('Error: $e');
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(
            content: Text('Erreur lors de la création de l\'applet')),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Create New Applet'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          children: [
            TextField(
              controller: ifController,
              decoration: const InputDecoration(
                labelText: 'If Condition',
              ),
            ),
            TextField(
              controller: thatController,
              decoration: const InputDecoration(
                labelText: 'That Condition',
              ),
            ),
            TextField(
              controller: ifTypeController,
              decoration: const InputDecoration(
                labelText: 'If Service',
              ),
            ),
            TextField(
              controller: thatTypeController,
              decoration: const InputDecoration(
                labelText: 'That Service',
              ),
            ),
            const SizedBox(height: 20),
            ElevatedButton(
              onPressed: () {
                postData();
              },
              child: const Text('Save Applet'),
            ),
          ],
        ),
      ),
    );
  }

  @override
  void dispose() {
    ifController.dispose();
    thatController.dispose();
    ifTypeController.dispose();
    thatTypeController.dispose();
    super.dispose();
  }
}

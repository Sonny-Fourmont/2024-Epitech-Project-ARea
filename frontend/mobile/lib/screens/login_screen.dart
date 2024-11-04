import 'package:area/screens/home_screen.dart';
import 'package:area/screens/webview_screen.dart';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:dio/dio.dart';

class LoginScreen extends StatefulWidget {
  const LoginScreen({Key? key}) : super(key: key);

  @override
  _LoginScreenState createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  final TextEditingController emailController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();
  final FlutterSecureStorage secureStorage = const FlutterSecureStorage();
  final Dio dio = Dio();

  String? errorMessage;
  bool isLoading = false;

  Future<void> login() async {
    setState(() {
      isLoading = true;
      errorMessage = null;
    });

    try {
      final response = await dio.post(
        dotenv.env['API_LOGIN_URL'] ?? 'http://localhost:8080/login',
        data: {
          'email': emailController.text,
          'password': passwordController.text
        },
      );
      await secureStorage.write(key: 'token', value: response.data['token']);

      if (mounted) {
        Navigator.of(context).pushReplacement(MaterialPageRoute(
            builder: (context) => HomeScreen(token: response.data['token'])));
      }
    } catch (e) {
      setState(() {
        errorMessage = e.toString();
        isLoading = false;
      });
    }
  }

  void googleLogin() {
    Navigator.push(
      context,
      MaterialPageRoute(
        builder: (context) => OAuthWebView(
          initialUrl: dotenv.env['GOOGLE_LOGIN_URL'] ?? 'https://localhost:8080/google/login',
        ),
      ),
    );
  }

  void githubLogin() {
    Navigator.push(
      context,
      MaterialPageRoute(
        builder: (context) => OAuthWebView(
          initialUrl: dotenv.env['GITHUB_LOGIN_URL'] ?? 'https://localhost:8080/github/login',
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Login'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          children: [
            TextField(
              controller: emailController,
              decoration: const InputDecoration(labelText: 'Email'),
            ),
            TextField(
              controller: passwordController,
              decoration: const InputDecoration(labelText: 'Password'),
              obscureText: true,
            ),
            const SizedBox(height: 20),
            if (errorMessage != null)
              Text(
                errorMessage!,
                style: const TextStyle(color: Colors.red),
              ),
            const SizedBox(height: 20),
            ElevatedButton(
              onPressed: isLoading ? null : login,
              child: isLoading
                  ? const CircularProgressIndicator(color: Colors.white)
                  : const Text('Login'),
            ),
            const SizedBox(height: 20),
            ElevatedButton.icon(
              icon: const Icon(Icons.login),
              onPressed: googleLogin,
              label: const Text('Login with Google'),
            ),
            const SizedBox(height: 20),
            ElevatedButton.icon(
              icon: const Icon(Icons.login),
              onPressed: githubLogin,
              label: const Text('Login with GitHub'),
            ),
          ],
        ),
      ),
    );
  }
}

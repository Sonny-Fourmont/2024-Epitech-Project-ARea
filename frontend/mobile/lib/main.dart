import 'package:area/screens/home_screen.dart';
import 'package:area/screens/login_screen.dart';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await dotenv.load();
  runApp(const MyApp());
}

class MyApp extends StatefulWidget {
  const MyApp({super.key});

  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  String? initialToken;
  bool isLoading = true;

  @override
  void initState() {
    super.initState();
    _loadInitialToken();
  }

  Future<void> _loadInitialToken() async {
    const secureStorage = FlutterSecureStorage();

    initialToken = await secureStorage.read(key: 'token');
    print("Initial token is $initialToken");

    setState(() {
      isLoading = false;
    });
  }

  @override
  Widget build(BuildContext context) {
    if (isLoading) {
      return const MaterialApp(
        debugShowCheckedModeBanner: false,
        home: Scaffold(
          body: Center(child: CircularProgressIndicator()),
        ),
      );
    }

    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Area',
      home: initialToken != null
          ? HomeScreen(token: initialToken)
          : const LoginScreen(),
    );
  }
}

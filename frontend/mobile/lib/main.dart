import 'package:area/bloc/auth_bloc.dart';
import 'package:area/bloc/auth_state.dart';
import 'package:area/screens/home_screen.dart';
import 'package:area/screens/login_screen.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
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
        home: Scaffold(
          body: Center(child: CircularProgressIndicator()),
        ),
      );
    }

    return MultiBlocProvider(
      providers: [
        BlocProvider(create: (_) => AuthBloc(initialToken)),
      ],
      child: MaterialApp(
        debugShowCheckedModeBanner: false,
        title: 'Area',
        routes: {
          '/login': (context) => const LoginScreen(),
          '/home': (context) => HomeScreen(token: initialToken),
        },
        home: BlocBuilder<AuthBloc, AuthState>(
          builder: (context, state) {
            if (state is AuthLoading) {
              return const Center(child: CircularProgressIndicator());
            }
            if (state is AuthSuccess) return HomeScreen(token: state.token);
            return const LoginScreen();
          },
        ),
      ),
    );
  }
}
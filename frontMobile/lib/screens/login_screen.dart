import 'package:area/bloc/auth_bloc.dart';
import 'package:area/bloc/auth_event.dart';
import 'package:area/bloc/auth_state.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class LoginScreen extends StatefulWidget {
  const LoginScreen({super.key});

  @override
  LoginScreenState createState() => LoginScreenState();
}

class LoginScreenState extends State<LoginScreen> {
  final TextEditingController emailController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Login'),
      ),
      body: BlocListener<AuthBloc, AuthState>(
        listener: (context, state) {
          if (state is AuthFailure) {
            ScaffoldMessenger.of(context)
                .showSnackBar(SnackBar(content: Text(state.error)));
          }

          // On navigue vers /home seulement si l'état est AuthSuccess
          if (state is AuthSuccess) {
            Navigator.pushReplacementNamed(context, '/home');
          }
        },
        child: BlocBuilder<AuthBloc, AuthState>(
          builder: (context, state) {
            if (state is AuthLoading) {
              return const Center(child: CircularProgressIndicator());
            }

            return Padding(
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
                  ElevatedButton(
                    onPressed: () {
                      BlocProvider.of<AuthBloc>(context).add(
                        LoginEvent(
                          emailController.text,
                          passwordController.text,
                        ),
                      );
                    },
                    child: const Text('Login'),
                  ),
                  const SizedBox(height: 20),
                  ElevatedButton.icon(
                    icon: const Icon(Icons.login),
                    onPressed: () {
                      BlocProvider.of<AuthBloc>(context)
                          .add(GoogleLoginEvent(context: context));
                    },
                    label: const Text('Google'),
                  ),
                  const SizedBox(height: 20),
                  ElevatedButton.icon(
                    icon: const Icon(Icons.login),
                    onPressed: () {
                      BlocProvider.of<AuthBloc>(context)
                          .add(GitHubLoginEvent(context: context));
                    },
                    label: const Text('GitHub'),
                  ),
                ],
              ),
            );
          },
        ),
      ),
    );
  }
}

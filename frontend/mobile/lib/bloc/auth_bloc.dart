import 'package:area/bloc/auth_event.dart';
import 'package:area/bloc/auth_state.dart';
import 'package:area/bloc/oauth_webview.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:dio/dio.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  final Dio dio = Dio();

  AuthBloc() : super(AuthInitial()) {
    on<LoginEvent>(_onLogin);
    on<RegisterEvent>(_onRegister);
    on<GoogleLoginEvent>(_onGoogleLogin);
    on<GitHubLoginEvent>(_onGitHubLogin);
  }
  void _onGoogleLogin(GoogleLoginEvent event, Emitter<AuthState> emit) async {
    emit(AuthLoading());
    try {
      final result = await Navigator.push(
        event.context,
        MaterialPageRoute(
          builder: (context) => const OAuthWebView(
            initialUrl: 'http://localhost:8080/google/login',
          ),
        ),
      );

      if (result != null && result['token'] != null) {
        emit(AuthSuccess(result['token']));
      } else {
        emit(AuthFailure('Google login failed'));
      }
    } catch (e) {
      emit(AuthFailure(e.toString()));
    }
  }

  void _onGitHubLogin(GitHubLoginEvent event, Emitter<AuthState> emit) async {
    emit(AuthLoading());
    try {
      final result = await Navigator.push(
        event.context,
        MaterialPageRoute(
          builder: (context) => const OAuthWebView(
            initialUrl: 'http://localhost:8080/github/login',
          ),
        ),
      );

      if (result != null && result['token'] != null) {
        emit(AuthSuccess(result['token']));
      } else {
        emit(AuthFailure('GitHub login failed'));
      }
    } catch (e) {
      emit(AuthFailure(e.toString()));
    }
  }

  void _onLogin(LoginEvent event, Emitter<AuthState> emit) async {
    emit(AuthLoading());
    try {
      final response = await dio.post(
        dotenv.env['API_LOGIN_URL']!,
        data: {'email': event.email, 'password': event.password},
      );
      emit(AuthSuccess(response.data['token']));
    } catch (e) {
      emit(AuthFailure(e.toString()));
    }
  }

  void _onRegister(RegisterEvent event, Emitter<AuthState> emit) async {
    emit(AuthLoading());
    try {
      final response = await dio.post(
        dotenv.env['API_REGISTER_URL']!,
        data: {'email': event.email, 'password': event.password},
      );
      emit(AuthSuccess(response.data['token']));
    } catch (e) {
      emit(AuthFailure(e.toString()));
    }
  }
}

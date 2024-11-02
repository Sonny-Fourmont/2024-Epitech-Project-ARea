import 'package:area/bloc/auth_event.dart';
import 'package:area/bloc/auth_state.dart';
import 'package:area/bloc/oauth_webview.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:dio/dio.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

//HOTFIX FOR ANDROID REPLACE localhost by localhost

class InitialTokenEvent extends AuthEvent {
  final String token;
  InitialTokenEvent(this.token);

}

void _onInitialToken(InitialTokenEvent event, Emitter<AuthState> emit, FlutterSecureStorage securestorage) async {
  emit(AuthLoading());
  try {
    securestorage.write(key: 'token', value: event.token);
    emit(AuthSuccess(event.token));
  } catch (e) {
    emit(AuthFailure(e.toString()));
  }
}

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  final Dio dio = Dio();
  final FlutterSecureStorage securestorage = const FlutterSecureStorage();

  AuthBloc(String? initialToken) : super(AuthInitial()) {
    on<LoginEvent>(_onLogin);
    on<RegisterEvent>(_onRegister);
    on<GoogleLoginEvent>(_onGoogleLogin);
    on<InitialTokenEvent>((event, emit) => _onInitialToken(event, emit, securestorage));
    if (initialToken != null) {
      print('Token found');
      add(InitialTokenEvent(initialToken));
    } else {
      print('No token');
    }
  }
  
  void _onGoogleLogin(GoogleLoginEvent event, Emitter<AuthState> emit) async {
    emit(AuthLoading());
    try {
      String urlServer = 'http://localhost:8080/google/login';
      // if (!) urlServer = 'http://localhost:8080/google/login';
      final result = await Navigator.push(
        event.context,
        MaterialPageRoute(
          builder: (context) => const OAuthWebView(
            initialUrl: 'http://localhost:8080/google/login',
          ),
        ),
      );

      if (result != null && result['token'] != null) {
        securestorage.write(key: 'token', value: result['token']);
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
        securestorage.write(key: 'token', value: result['token']);
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
        '/users/login',
        data: {'email': event.email, 'password': event.password},
      );
      securestorage.write(key: 'token', value: response.data['token']);
      emit(AuthSuccess(response.data['token']));
    } catch (e) {
      emit(AuthFailure(e.toString()));
    }
  }

  void _onRegister(RegisterEvent event, Emitter<AuthState> emit) async {
    emit(AuthLoading());
    try {
      final response = await dio.post(
        '/users/register',
        data: {'email': event.email, 'password': event.password},
      );
      securestorage.write(key: 'token', value: response.data['token']);
      emit(AuthSuccess(response.data['token']));
    } catch (e) {
      emit(AuthFailure(e.toString()));
    }
  }
}

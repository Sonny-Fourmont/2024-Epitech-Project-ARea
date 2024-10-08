import 'package:area/bloc/auth_event.dart';
import 'package:area/bloc/auth_state.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:dio/dio.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:url_launcher/url_launcher_string.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  final Dio dio = Dio();

  AuthBloc() : super(AuthInitial()) {
    on<LoginEvent>(_onLogin);
    on<RegisterEvent>(_onRegister);
    on<GoogleLoginEvent>(_onGoogleLogin);
    on<GithubLoginEvent>(_onGithubLogin);
  }

  void _onGoogleLogin(GoogleLoginEvent event, Emitter<AuthState> emit) async {
    emit(AuthLoading());
    final url = dotenv.env['GOOGLE_LOGIN_URL']!;

    if (await canLaunchUrlString(url)) {
      await launchUrlString(url);
      emit(AuthSuccess("La page de connexion Google a été ouverte."));
    } else {
      emit(AuthFailure("Impossible d'ouvrir le lien de connexion Google."));
    }
  }

  void _onGithubLogin(GithubLoginEvent event, Emitter<AuthState> emit) async {
    emit(AuthLoading());
    final url = dotenv.env['GITHUB_LOGIN_URL']!;

    if (await canLaunchUrlString(url)) {
      await launchUrlString(url);
      emit(AuthSuccess("La page de connexion GitHub a été ouverte."));
    } else {
      emit(AuthFailure("Impossible d'ouvrir le lien de connexion GitHub."));
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

import 'package:equatable/equatable.dart';
import 'package:flutter/material.dart';

abstract class AuthEvent extends Equatable {
  @override
  List<Object?> get props => [];
}

class LoginEvent extends AuthEvent {
  final String email;
  final String password;

  LoginEvent(this.email, this.password);

  @override
  List<Object?> get props => [email, password];
}

class RegisterEvent extends AuthEvent {
  final String email;
  final String password;

  RegisterEvent(this.email, this.password);

  @override
  List<Object?> get props => [email, password];
}

class GoogleLoginEvent extends AuthEvent {
  final BuildContext context;

  GoogleLoginEvent({required this.context});

  @override
  List<Object?> get props => [context];
}

class GitHubLoginEvent extends AuthEvent {
  final BuildContext context;

  GitHubLoginEvent({required this.context});

  @override
  List<Object?> get props => [context];
}

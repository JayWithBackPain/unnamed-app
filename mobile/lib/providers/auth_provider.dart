import 'package:flutter/material.dart';
import '../services/api_service.dart';

class AuthProvider with ChangeNotifier {
  final ApiService _apiService = ApiService();
  bool _isAuthenticated = false;
  Map<String, dynamic>? _user;

  bool get isAuthenticated => _isAuthenticated;
  Map<String, dynamic>? get user => _user;

  // 登入
  Future<void> login(String username, String password) async {
    try {
      final response = await _apiService.login(
        username: username,
        password: password,
      );
      _user = response['user'];
      _isAuthenticated = true;
      notifyListeners();
    } catch (e) {
      rethrow;
    }
  }

  // 註冊
  Future<void> register({
    required String username,
    required String password,
    required String email,
    required DateTime birthDate,
  }) async {
    try {
      await _apiService.register(
        username: username,
        password: password,
        email: email,
        birthDate: birthDate,
      );
    } catch (e) {
      rethrow;
    }
  }

  // 登出
  Future<void> logout() async {
    _isAuthenticated = false;
    _user = null;
    notifyListeners();
  }
} 
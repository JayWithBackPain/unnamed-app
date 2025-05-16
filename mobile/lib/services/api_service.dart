import 'dart:convert';
import 'package:dio/dio.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class ApiService {
  final Dio _dio = Dio();
  final FlutterSecureStorage _storage = const FlutterSecureStorage();
  static const String baseUrl = 'http://localhost:8080/api';

  ApiService() {
    _dio.interceptors.add(InterceptorsWrapper(
      onRequest: (options, handler) async {
        final token = await _storage.read(key: 'token');
        if (token != null) {
          options.headers['Authorization'] = 'Bearer $token';
        }
        return handler.next(options);
      },
    ));
  }

  // 用戶註冊
  Future<Map<String, dynamic>> register({
    required String username,
    required String password,
    required String email,
    required DateTime birthDate,
  }) async {
    try {
      final response = await _dio.post(
        '$baseUrl/register',
        data: {
          'username': username,
          'password': password,
          'email': email,
          'birth_date': birthDate.toIso8601String(),
        },
      );
      return response.data;
    } catch (e) {
      throw _handleError(e);
    }
  }

  // 用戶登入
  Future<Map<String, dynamic>> login({
    required String username,
    required String password,
  }) async {
    try {
      final response = await _dio.post(
        '$baseUrl/login',
        data: {
          'username': username,
          'password': password,
        },
      );
      if (response.data['token'] != null) {
        await _storage.write(key: 'token', value: response.data['token']);
      }
      return response.data;
    } catch (e) {
      throw _handleError(e);
    }
  }

  // 生成奇門盤
  Future<Map<String, dynamic>> generateChart(DateTime dateTime) async {
    try {
      final response = await _dio.post(
        '$baseUrl/qimen/chart',
        data: {
          'date_time': dateTime.toIso8601String(),
        },
      );
      return response.data;
    } catch (e) {
      throw _handleError(e);
    }
  }

  // 錯誤處理
  Exception _handleError(dynamic error) {
    if (error is DioException) {
      if (error.response?.data != null) {
        return Exception(error.response?.data['error'] ?? '未知錯誤');
      }
      return Exception(error.message ?? '網絡錯誤');
    }
    return Exception('未知錯誤');
  }
} 
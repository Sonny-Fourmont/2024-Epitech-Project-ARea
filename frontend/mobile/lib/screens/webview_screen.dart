import 'dart:io';
import 'package:area/screens/home_screen.dart';
import 'package:area/screens/login_screen.dart';
import 'package:flutter/material.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:webview_flutter/webview_flutter.dart';
import 'package:webview_flutter_android/webview_flutter_android.dart';
import 'package:webview_flutter_wkwebview/webview_flutter_wkwebview.dart';

class OAuthWebView extends StatefulWidget {
  final String initialUrl;

  const OAuthWebView({super.key, required this.initialUrl});

  @override
  OAuthWebViewState createState() => OAuthWebViewState();
}

class OAuthWebViewState extends State<OAuthWebView> {
  late final WebViewController _controller;
  final FlutterSecureStorage secureStorage = const FlutterSecureStorage();
  @override
  void initState() {
    super.initState();

    // Configure the WebView for Android or iOS
    late final PlatformWebViewControllerCreationParams params;
    if (WebViewPlatform.instance is WebKitWebViewPlatform) {
      params = WebKitWebViewControllerCreationParams(
        allowsInlineMediaPlayback: true,
        mediaTypesRequiringUserAction: const <PlaybackMediaTypes>{},
      );
    } else {
      params = const PlatformWebViewControllerCreationParams();
    }

    _controller = WebViewController.fromPlatformCreationParams(params);
    String userAgent = Platform.isAndroid
        ? 'Mozilla/5.0 (Linux; Android 15) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.81 Mobile Safari/537.36'
        : 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 Safari/605.1.15';

    _controller.setUserAgent(userAgent);

    _controller
      ..setJavaScriptMode(JavaScriptMode.unrestricted)
      ..setNavigationDelegate(
        NavigationDelegate(
          onPageStarted: (String url) {
            debugPrint('Page started loading: $url');
            if (Platform.isAndroid && url.contains('localhost')) {
              url = url.replaceAll('localhost', '10.0.2.2');
              _controller.loadRequest(Uri.parse(url));
              return;
            }
          },
          onNavigationRequest: (NavigationRequest request) async {
            debugPrint('Navigating to: ${request.url}');
            if (request.url.contains('3000/home')) {
              final tokenPattern = RegExp(r'token=([^&]+)');
              final tokenMatch = tokenPattern.firstMatch(request.url);
              final String token = tokenMatch?.group(1) ?? "";
              debugPrint("token : $token");

              if (token.isNotEmpty) {
                await secureStorage.write(key: 'token', value: token);
                if (mounted) {
                  Navigator.of(context).pushReplacement(MaterialPageRoute(
                      builder: (context) => HomeScreen(token: token)));
                }
              } else {
                Navigator.of(context).pushReplacement(MaterialPageRoute(
                    builder: (context) => const LoginScreen()));
              }
              return NavigationDecision.prevent;
            }
            return NavigationDecision.navigate;
          },
        ),
      )
      ..loadRequest(Uri.parse(widget.initialUrl));

    if (_controller.platform is AndroidWebViewController) {
      AndroidWebViewController.enableDebugging(true);
      (_controller.platform as AndroidWebViewController)
          .setMediaPlaybackRequiresUserGesture(false);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('OAuth Login'),
      ),
      body: WebViewWidget(controller: _controller),
    );
  }
}

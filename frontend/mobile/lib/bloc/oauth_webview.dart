import 'dart:io';
import 'dart:math';

import 'package:flutter/material.dart';
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

  String cleanJson(String jsonString) {
    String cleanedString = jsonString.replaceAll(r'\', '');
    return cleanedString;
  }

  @override
  void initState() {
    super.initState();

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
            debugPrint('Params: $params');
          },
          onPageFinished: (String url) async {
            debugPrint('Page finished loading: $url');

            if (url.contains('google/?state=state-token')) {
              Object response = await _controller
                  .runJavaScriptReturningResult('document.body.innerText');
              String resStr = response.toString();
              debugPrint("Response: $resStr");
              if (context.mounted) {
                if (response != "") {
                  resStr = cleanJson(resStr);
                  String token = resStr.substring(
                      resStr.indexOf('token') + 8, resStr.indexOf('"}'));
                  // ignore: use_build_context_synchronously
                  Navigator.pop(context, {'token': token});
                } else {
                  // ignore: use_build_context_synchronously
                  Navigator.pop(context);
                }
              }
            }
          },
          onNavigationRequest: (NavigationRequest request) async {
            debugPrint('Navigating to: ${request.url}');
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

import 'dart:convert';
import 'dart:io';

import 'package:area/component/appBar.dart';
import 'package:area/component/navBar.dart';
import 'package:area/screens/home_screen.dart';
import 'package:area/screens/login_screen.dart';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class IfThat {
  final String type;
  final List<String> options;
  final String prettyName;
  final String description;
  final String tokenName;
  final String urlLogin;

  IfThat({
    required this.type,
    required this.options,
    required this.prettyName,
    required this.description,
    required this.tokenName,
    required this.urlLogin,
  });

  factory IfThat.fromJson(Map<String, dynamic> json) {
    return IfThat(
      type: json['Type'] as String,
      options: List<String>.from(json['Options'] ?? []),
      prettyName: json['PrettyName'] as String,
      description: json['Description'] as String,
      tokenName: json['TokenName'] as String,
      urlLogin: json['UrlLogin'] as String,
    );
  }
}

class ServiceAvailable {
  final List<IfThat> ifList;
  final List<IfThat> thatList;

  ServiceAvailable({
    required this.ifList,
    required this.thatList,
  });

  factory ServiceAvailable.fromJson(Map<String, dynamic> json) {
    return ServiceAvailable(
      ifList: (json['If'] as List<dynamic>)
          .map((item) => IfThat.fromJson(item))
          .toList(),
      thatList: (json['That'] as List<dynamic>)
          .map((item) => IfThat.fromJson(item))
          .toList(),
    );
  }
}

class NewAppletScreen extends StatefulWidget {
  final String token;
  
  const NewAppletScreen({super.key, required this.token});
  @override
  NewAppletScreenState createState() => NewAppletScreenState();
}

class NewAppletScreenState extends State<NewAppletScreen> {
  final Dio dio = Dio();
  List<IfThat> ifOptions = [];
  List<IfThat> thatOptions = [];
  IfThat? selectedIf;
  IfThat? selectedThat;

  final TextEditingController ifOptionController = TextEditingController();
  final TextEditingController thatOptionController = TextEditingController();

  @override
  void initState() {
    super.initState();
    fetchData();
  }

  Future<void> fetchData() async {
    final apiUrl = dotenv.env['API_SERVICE_URL'] ?? 'http://localhost:8080/services';
    try {
      final response = await dio.get(apiUrl);

      final Map<String, dynamic> data = response.data is String
          ? jsonDecode(response.data)
          : response.data;  

      final serviceData = ServiceAvailable.fromJson(data);  

      setState(() {
        ifOptions = serviceData.ifList;
        thatOptions = serviceData.thatList;
      });
    } catch (e) {
      // ignore: use_build_context_synchronously
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Erreur lors du chargement des données')),
      );
    }
  }

  void postData() async {
    if (selectedIf == null || selectedThat == null) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Veuillez sélectionner un If et un That')),
      );
      return;
    }

    var apiUrl = dotenv.env['API_APPLET_URL'] ?? 'http://localhost:8080/applet/';
    if (Platform.isAndroid) {
      apiUrl = apiUrl.replaceAll('localhost', '10.0.2.2');
    }
    try {
      final data = {
        "is_on": true,
        "if": selectedIf!.options.isNotEmpty ? ifOptionController.text : '',
        "that": selectedThat!.options.isNotEmpty ? thatOptionController.text : '',
        "if_type": selectedIf!.type,
        "that_type": selectedThat!.type,
      };
      dio.options.headers['Authorization'] = 'Bearer ${widget.token}';
      await dio.post(apiUrl, data: data);
      // ignore: use_build_context_synchronously
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Applet créé avec succès !')),
      );
    } catch (e) {
      // ignore: use_build_context_synchronously
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Erreur lors de la création de l\'applet')),
      );
    }
  }

  Widget buildServiceList(List<IfThat> options, String type) {
    return Column(
      children: options.map((service) {
        final isSelected =
            (type == "If" && selectedIf == service) ||
            (type == "That" && selectedThat == service);

        return ListTile(
          title: Text(service.prettyName, style: const TextStyle(fontWeight: FontWeight.bold)),
          subtitle: Text(service.description, style: const TextStyle(fontSize: 12)),
          selected: isSelected,
          selectedTileColor: isSelected ? Colors.black87 : Colors.transparent,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(10),
          ),
          selectedColor: Colors.white,
          textColor: isSelected ? Colors.white : Colors.black,
          onTap: () {
            setState(() {
              if (type == "If") {
                selectedIf = service;
                ifOptionController.clear();
              } else {
                selectedThat = service;
                thatOptionController.clear();
              }
            });
          },
        );
      }).toList(),
    );
  }
  void disconnection() {
    const securestorage = FlutterSecureStorage();
    securestorage.delete(key: 'token');
    Navigator.push(
      context,
      MaterialPageRoute(
        builder: (context) => const LoginScreen(),
      ),
    );
  }

  void disconnect() {
    showDialog(
      context: context,
      builder: (context) {
        return AlertDialog(
          title: const Text("Disconnect"),
          content: const Text("Are you sure you want to disconnect?"),
          actions: [
            TextButton(
              onPressed: () => Navigator.of(context).pop(),
              child: const Text('Cancel'),
            ),
            TextButton(
              onPressed: () {
                Navigator.of(context).pop();
                disconnection();
              },
              child: const Text('Disconnect'),
            ),
          ],
        );
      },
    );
  }
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const CustomAppBar(),
      bottomNavigationBar: CustomNavBar(
        currentIndex: 1,
        onTap: (index) {
          if (index == 0) {
             Navigator.pushReplacement(
              context,
              MaterialPageRoute(
                builder: (context) => HomeScreen(token: widget.token),
              ),
            );
          }  else if (index == 2) {
            disconnect();
          }

        },
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            const Text(
              'Select "If" Service',
              style: TextStyle(fontWeight: FontWeight.bold),
            ),
            SizedBox(
              height: 200,
              child: SingleChildScrollView(
                child: buildServiceList(ifOptions, "If"),
              ),
            ),


            if (selectedIf != null && selectedIf!.options.isNotEmpty) ...[
              Container(
                margin: const EdgeInsets.only(top: 20),
                decoration: BoxDecoration(
                  color: Colors.black87,
                  borderRadius: BorderRadius.circular(10),
                ),
                child:
              TextField(
                controller: ifOptionController,
                cursorColor: Colors.white,
                style: const TextStyle(color: Colors.white),
                decoration: InputDecoration(
                  contentPadding: const EdgeInsets.all(10),
                  border: InputBorder.none,
                  labelText: '> Please enter "${selectedIf!.options[0]}"',
                  labelStyle: const TextStyle(color: Colors.white),
                
                ),
              ),),
            ],

            const Divider(),
            const Text(
              'Select "That" Service',
              style: TextStyle(fontWeight: FontWeight.bold),
            ),
            SizedBox(
              height: 200,
              child: SingleChildScrollView(
                child: buildServiceList(thatOptions, "That"),
              ),
            ),

            if (selectedThat != null && selectedThat!.options.isNotEmpty) ...[
              Container(
                margin: const EdgeInsets.only(top: 20),
                decoration: BoxDecoration(
                  color: Colors.black87,
                  borderRadius: BorderRadius.circular(10),
                ),
                child:
              TextField(
                controller: thatOptionController,
                cursorColor: Colors.white,
                style: const TextStyle(color: Colors.white),
                decoration: InputDecoration(
                  contentPadding: const EdgeInsets.all(10),
                  border: InputBorder.none,
                  labelText: '> Please enter "${selectedThat!.options[0]}"',
                  labelStyle: const TextStyle(color: Colors.white),
                
                ),
              ),
              ),
              
              
            ],

            const SizedBox(height: 20),
            ElevatedButton(
              onPressed: selectedIf != null && selectedThat != null ? postData : null,
              style: ElevatedButton.styleFrom(
                backgroundColor: Colors.black87,
              ),
              child: const Text('Save Applet', style: TextStyle(color: Colors.white)),
            ),
          ],
        ),
      ),

    );
  }


  @override
  void dispose() {
    ifOptionController.dispose();
    thatOptionController.dispose();
    super.dispose();
  }
}

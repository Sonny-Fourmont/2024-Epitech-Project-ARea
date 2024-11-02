import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:http/http.dart' as http;
import '../component/appBar.dart';
import '../component/navBar.dart';
import '../component/searchBar.dart'; 


class Applet {
  final String id;
  final String idUser;
  final DateTime createdAt;
  final DateTime updatedAt;
  final bool isOn;
  final String ifCondition;
  final String thatCondition;
  final String ifType;
  final String thatType;

  Applet({
    required this.id,
    required this.idUser,
    required this.createdAt,
    required this.updatedAt,
    required this.isOn,
    required this.ifCondition,
    required this.thatCondition,
    required this.ifType,
    required this.thatType,
  });

  factory Applet.fromJson(Map<String, dynamic> json) {
    return Applet(
      id: json['ID'],
      idUser: json['ID_User'],
      createdAt: DateTime.parse(json['CreatedAt']),
      updatedAt: DateTime.parse(json['UpdatedAt']),
      isOn: json['IsOn'],
      ifCondition: json['If'],
      thatCondition: json['That'],
      ifType: json['IfType'],
      thatType: json['ThatType'],
    );
  }
}

class HomeScreen extends StatefulWidget {
  final String? token;

  const HomeScreen({super.key, required this.token});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  late Future<List<Applet>> futureApplets;

  @override
  void initState() {
    super.initState();
    futureApplets = fetchApplets();
  }

  @override
  void didUpdateWidget(HomeScreen oldWidget) {
    super.didUpdateWidget(oldWidget);
    futureApplets = fetchApplets();
  }

  Future<List<Applet>> fetchApplets() async {
    final apiAppletUrl = dotenv.env['API_APPLET_URL']!;
    final token = widget.token;
    final response = await http.get(
      Uri.parse(apiAppletUrl),
      headers: {
        'Authorization': 'Bearer $token',
      },
    );
  
    if (response.statusCode == 200) {
      try {
        final decodedJson = json.decode(response.body);
        if (decodedJson is String) {
          final jsonString = decodedJson;
          final Map<String, dynamic> jsonMap = json.decode(jsonString);

          final List<dynamic> appletArray = jsonMap['applet_array'];

          return appletArray.map((appletJson) => Applet.fromJson(appletJson)).toList();
        } else if (decodedJson is Map<String, dynamic>) {
          final List<dynamic> appletArray = decodedJson['applet_array'];

          return appletArray.map((appletJson) => Applet.fromJson(appletJson)).toList();
        } else {
          return [];
        }
      } catch (e) {
        return [];
      }
    } else {
      throw Exception("Failed to load applets");
    }
}

  Widget buildAppletsList(List<Applet> applets) {
    return ListView.builder(
      itemCount: applets.length,
      shrinkWrap: true,
      itemBuilder: (context, index) {
        final applet = applets[index];
        return Card(
          margin: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
          elevation: 4,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(12),
          ),
          child: ListTile(
            title: Text(
              "ACTION: ${applet.ifCondition}",
              style: const TextStyle(fontWeight: FontWeight.bold),
            ),
            subtitle: Text("REACTION: ${applet.thatCondition}"),
            trailing: Icon(
              applet.isOn ? Icons.check_circle : Icons.cancel,
              color: applet.isOn ? Colors.green : Colors.red,
            ),
            onTap: () => showAppletDetails(applet),
          ),
        );
      },
    );
  }

  void showAppletDetails(Applet applet) {
    showDialog(
      context: context,
      builder: (context) {
        return AlertDialog(
          title: const Text("More Information"),
          content: Column(
            mainAxisSize: MainAxisSize.min,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text("User ID: ${applet.idUser}"),
              Text("If Type: ${applet.ifType}"),
              Text("That Type: ${applet.thatType}"),
              Text("Created At: ${applet.createdAt}"),
              Text("Updated At: ${applet.updatedAt}"),
              Text("Status: ${applet.isOn ? 'On' : 'Off'}"),
            ],
          ),
          actions: [
            TextButton(
              onPressed: () => Navigator.of(context).pop(),
              child: const Text('Close'),
            ),
          ],
        );
      },
    );
  }

  void disconnection() {
    const securestorage = FlutterSecureStorage();
    securestorage.delete(key: 'token');
    Navigator.of(context).restorablePushReplacementNamed('/login');
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
        currentIndex: 0,
        onTap: (index) {
          if (index == 2) {
            disconnect();
          }
          print("Tapped index: $index");
        },
      ),
      body: Container(
        padding: const EdgeInsets.all(16),
        child: Column(
          children: [
            const Text(
              'Explore',
              style: TextStyle(fontSize: 30, fontWeight: FontWeight.bold),
            ),
            const SizedBox(height: 10),
            const Text(
              'Automate to save time and get more done',
              style: TextStyle(fontSize: 15, fontWeight: FontWeight.bold),
              textAlign: TextAlign.center,
            ),
            const SizedBox(height: 40),
            Expanded(
              child: FutureBuilder<List<Applet>>(
                future: futureApplets,
                builder: (context, snapshot) {
                  if (snapshot.connectionState == ConnectionState.waiting) {
                    return const Center(child: CircularProgressIndicator());
                  } else if (snapshot.hasError) {
                    return Center(child: Text("Erreur : ${snapshot.error}"));
                  } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
                    return const Center(child: Text("Aucun applet trouvé"));
                  } else {
                    return buildAppletsList(snapshot.data!);
                  }
                },
              ),
            ),
          ],
        ),
      ),
    );
  }
}



class EntrySearch extends StatelessWidget {
  const EntrySearch({Key? key, required this.onSearchEntered})
      : super(key: key);

  final Function(String) onSearchEntered;
  @override
  Widget build(BuildContext context) {
    var padinputText = 14.0;
    return Padding(
      padding: const EdgeInsets.only(top: 112),
      child: Align(
        alignment: Alignment.topCenter,
        child: SizedBox(
          height: 48,
          width: 350,
          child: Align(
            alignment: Alignment.topLeft,
            child: SizedBox(
              height: 48,
              width: 331,
              child: Stack(
                children: [
                  Container(
                    width: double.maxFinite,
                    height: 48,
                    decoration: BoxDecoration(
                      color: const Color.fromARGB(255, 235, 233, 229),
                      borderRadius: BorderRadius.circular(8),
                    ),
                    child: TextFormField(
                      style: const TextStyle(
                        fontSize: 16,
                        fontFamily: 'Avenir Next',
                        height: 1.25,
                        color: Color.fromARGB(255, 43, 42, 40),
                        fontWeight: FontWeight.w500,
                      ),
                      cursorColor: const Color.fromARGB(255, 254, 152, 97),
                      decoration: InputDecoration(
                        // hintText: AppLocalizations.of(context).docs_search,
                        hintStyle: const TextStyle(
                          fontSize: 16,
                          height: 1.25,
                          fontFamily: 'Avenir Next',
                          color: Color.fromARGB(255, 163, 159, 166),
                          fontWeight: FontWeight.w500,
                        ),
                        border: InputBorder.none,
                        isCollapsed: true,
                        contentPadding:
                            EdgeInsets.fromLTRB(44, padinputText, 0, 0),
                      ),
                      onChanged: (value) {
                        onSearchEntered(value);
                      },
                    ),
                  ),
                  const Padding(
                    padding: EdgeInsets.only(left: 12),
                    child: Align(
                      alignment: Alignment.centerLeft,
                      // child: Image.asset(
                      //   iconLoop,
                      //   height: 24,
                      //   width: 24,
                      // ),
                    ),
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
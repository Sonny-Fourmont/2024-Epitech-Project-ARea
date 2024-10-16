import 'package:flutter/material.dart';

class customSearchBar extends StatelessWidget {
  const customSearchBar({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: TextField(
        decoration: InputDecoration(
          hintText: 'Rechercher...',
          prefixIcon: const Icon(Icons.search),
          border: OutlineInputBorder(
            borderRadius: BorderRadius.circular(12),
            borderSide: BorderSide.none,
          ),
          filled: true,  
          fillColor: Colors.grey[200],
        ),
        onChanged: (value) {
          print('Recherche : $value');
        },
      ),
    );
  }
}

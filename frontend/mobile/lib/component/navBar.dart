import 'package:flutter/material.dart';

class CustomNavBar extends StatelessWidget {
  final int currentIndex;
  final ValueChanged<int> onTap;

  const CustomNavBar({
    super.key,
    required this.currentIndex,
    required this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return BottomNavigationBar(
      items: const <BottomNavigationBarItem>[
        BottomNavigationBarItem(
          icon: Icon(Icons.insert_drive_file, size: 30),
          label: 'Create Applet',
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.calendar_today, size: 30),
          label: 'My Applet',
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.account_circle, size: 30),
          label: 'Account',
        ),
      ],
      currentIndex: currentIndex,
      selectedItemColor: Colors.black,
      unselectedItemColor: Colors.grey,
      onTap: onTap,
    );
  }
}

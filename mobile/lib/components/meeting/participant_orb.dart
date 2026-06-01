import 'dart:math';
import 'package:flutter/material.dart';
import '../../components/shared/app_theme.dart';
import '../../controllers/room_controller.dart';

class ParticipantOrb extends StatefulWidget {
  final bool isActive;
  final String participantName;
  final String userID;
  final RoomController roomController;
  final double size;

  const ParticipantOrb({
    super.key,
    required this.isActive,
    required this.participantName,
    required this.userID,
    required this.roomController,
    this.size = 120,
  });

  @override
  State<ParticipantOrb> createState() => _ParticipantOrbState();
}

class _ParticipantOrbState extends State<ParticipantOrb>
    with SingleTickerProviderStateMixin {
  late AnimationController _controller;

  @override
  void initState() {
    super.initState();
    _controller = AnimationController(
      vsync: this,
      duration: const Duration(seconds: 4),
    )..repeat(reverse: true);
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onLongPress: () {
        showDialog(
          context: context,
          builder: (context) => AlertDialog(
            title: const Text('Report Participant'),
            content: const Text(
              'Are you sure you want to report this participant\'s segment?',
            ),
            actions: [
              TextButton(
                onPressed: () => Navigator.pop(context),
                child: const Text('Cancel'),
              ),
              TextButton(
                onPressed: () {
                  // Report segment (passing 1 for roundIndex as placeholder)
                  widget.roomController.reportSegment(widget.userID, 1);
                  print('Reporting ${widget.userID}');
                  Navigator.pop(context);
                },
                child: const Text('Report'),
              ),
            ],
          ),
        );
      },
      child: Column(
        children: [
          AnimatedBuilder(
            animation: _controller,
            builder: (context, child) {
              final breathing = 0.8 + 0.2 * _controller.value;
              return Container(
                width: widget.size * breathing,
                height: widget.size * breathing,
                decoration: BoxDecoration(
                  color: widget.isActive
                      ? RifqahColors.primary
                      : RifqahColors.secondaryContainer,
                  shape: BoxShape.circle,
                  boxShadow: widget.isActive
                      ? [
                          BoxShadow(
                            color: RifqahColors.primary.withOpacity(0.3),
                            blurRadius: 30 * breathing,
                            spreadRadius: 5 * breathing,
                          ),
                        ]
                      : [],
                ),
                child: Center(
                  child: Icon(
                    Icons.waves,
                    color: Colors.white.withOpacity(0.4),
                    size: widget.size * 0.4,
                  ),
                ),
              );
            },
          ),
          const SizedBox(height: 12),
          Text(
            widget.participantName,
            style: Theme.of(context).textTheme.labelSmall,
          ),
        ],
      ),
    );
  }
}

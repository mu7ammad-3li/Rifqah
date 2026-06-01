import 'package:flutter/material.dart';
import '../../controllers/audio_controller.dart';
import '../shared/app_theme.dart';

class TapToTalkButton extends StatefulWidget {
  final AudioController audioController;
  final bool isMyTurn;

  const TapToTalkButton({
    super.key,
    required this.audioController,
    required this.isMyTurn,
  });

  @override
  State<TapToTalkButton> createState() => _TapToTalkButtonState();
}

class _TapToTalkButtonState extends State<TapToTalkButton> {
  bool _isCapturing = false;

  void _toggleCapture() {
    setState(() {
      _isCapturing = !_isCapturing;
    });

    if (_isCapturing) {
      widget.audioController.startCapture();
    } else {
      widget.audioController.stopCapture();
    }
  }

  @override
  Widget build(BuildContext context) {
    // Only allow interaction if it's the user's turn
    final canInteract = widget.isMyTurn;
    final color = _isCapturing
        ? RifqahColors.secondary
        : (canInteract ? RifqahColors.primary : Colors.grey);

    return GestureDetector(
      onTap: canInteract ? _toggleCapture : null,
      child: Container(
        width: 100,
        height: 100,
        decoration: BoxDecoration(color: color, shape: BoxShape.circle),
        child: Center(
          child: Icon(
            _isCapturing ? Icons.mic : Icons.mic_none,
            color: Colors.white,
            size: 48,
          ),
        ),
      ),
    );
  }
}

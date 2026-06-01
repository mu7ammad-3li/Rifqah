import 'dart:convert';
import 'dart:math';
import 'package:lazysodium/lazysodium.dart';

void main() async {
  final sodium = Lazysodium.instance();
  // We need to initialize the library properly
  await sodium.ready;
  
  final kp = sodium.cryptoBoxKeypair();
  
  // Use standard dart:convert for base64
  print('Public Key (Base64): ' + base64.encode(kp.publicKey));
  print('Private Key (Base64): ' + base64.encode(kp.secretKey));
}

import 'package:flutter/material.dart';
import '../../components/shared/app_theme.dart';
import '../meeting/meeting_room_screen.dart';
import 'package:uuid/uuid.dart';

class UnauthenticatedDiscoveryScreen extends StatelessWidget {
  const UnauthenticatedDiscoveryScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: RifqahColors.surface,
        elevation: 0,
        centerTitle: true,
        title: Row(
          mainAxisSize: MainAxisSize.min,
          children: [
            const Icon(Icons.shield_outlined, color: RifqahColors.primary),
            const SizedBox(width: RifqahSpacing.base),
            Text(
              'Rifqah (رِفْقَة)',
              style: Theme.of(
                context,
              ).textTheme.headlineMedium?.copyWith(fontSize: 20),
            ),
          ],
        ),
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.symmetric(
          horizontal: RifqahSpacing.containerPaddingMobile,
        ),
        child: Column(
          children: [
            const SizedBox(height: RifqahSpacing.base),
            _buildSearchField(),
            const SizedBox(height: RifqahSpacing.sectionGap),
            _buildExploreCircles(context),
            const SizedBox(height: RifqahSpacing.sectionGap),
          ],
        ),
      ),
    );
  }

  Widget _buildSearchField() {
    return Container(
      decoration: BoxDecoration(
        color: RifqahColors.tertiaryFixedDim.withOpacity(0.2),
        borderRadius: RifqahShapes.full,
      ),
      child: const TextField(
        decoration: InputDecoration(
          hintText: 'Search for a support circle or Room ID...',
          prefixIcon: Icon(Icons.search, color: RifqahColors.outline),
          border: InputBorder.none,
          contentPadding: EdgeInsets.symmetric(vertical: 20, horizontal: 20),
        ),
      ),
    );
  }

  Widget _buildExploreCircles(BuildContext context) {
    return Column(
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Text(
              'Explore Circles',
              style: Theme.of(context).textTheme.headlineMedium,
            ),
            TextButton(onPressed: () {}, child: const Text('View All')),
          ],
        ),
        const SizedBox(height: RifqahSpacing.stackGap),
        _buildCircleCard(
          context,
          title: 'Calm Horizon',
          subtitle: 'Navigating Anxiety Together',
          description:
              'A gentle space to share experiences and grounding techniques.',
        ),
      ],
    );
  }

  Widget _buildCircleCard(
    BuildContext context, {
    required String title,
    required String subtitle,
    required String description,
  }) {
    return Container(
      padding: const EdgeInsets.all(24),
      decoration: const BoxDecoration(
        color: Color(0xFFF4EBD0),
        borderRadius: RifqahShapes.md,
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            children: [
              Container(
                width: 56,
                height: 56,
                decoration: const BoxDecoration(
                  color: RifqahColors.primaryContainer,
                  shape: BoxShape.circle,
                ),
              ),
              const SizedBox(width: 16),
              Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(title, style: Theme.of(context).textTheme.labelMedium),
                  Text(
                    subtitle,
                    style: Theme.of(context).textTheme.bodyMedium?.copyWith(
                      fontWeight: FontWeight.w600,
                    ),
                  ),
                ],
              ),
            ],
          ),
          const SizedBox(height: 24),
          Text(description, style: Theme.of(context).textTheme.bodyMedium),
          const SizedBox(height: 24),
          SizedBox(
            width: double.infinity,
            child: ElevatedButton(
              onPressed: () {
                // TODO: Implement proper room joining logic with backend
              },
              style: ElevatedButton.styleFrom(
                backgroundColor: RifqahColors.secondaryContainer,
                shape: const RoundedRectangleBorder(
                  borderRadius: RifqahShapes.full,
                ),
              ),
              child: const Text('Join Room'),
            ),
          ),
        ],
      ),
    );
  }
}

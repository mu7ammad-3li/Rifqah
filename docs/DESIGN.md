---
name: Serene Sanctuary
colors:
  surface: '#fbf9f5'
  surface-dim: '#dbdad6'
  surface-bright: '#fbf9f5'
  surface-container-lowest: '#ffffff'
  surface-container-low: '#f5f3ef'
  surface-container: '#efeeea'
  surface-container-high: '#eae8e4'
  surface-container-highest: '#e4e2de'
  on-surface: '#1b1c1a'
  on-surface-variant: '#414848'
  inverse-surface: '#30312e'
  inverse-on-surface: '#f2f0ed'
  outline: '#717878'
  outline-variant: '#c1c8c7'
  surface-tint: '#476363'
  primary: '#062425'
  on-primary: '#ffffff'
  primary-container: '#1e3a3a'
  on-primary-container: '#86a4a3'
  inverse-primary: '#aecccc'
  secondary: '#7b563f'
  on-secondary: '#ffffff'
  secondary-container: '#fecdb0'
  on-secondary-container: '#79553e'
  tertiary: '#24200f'
  on-tertiary: '#ffffff'
  tertiary-container: '#3a3523'
  on-tertiary-container: '#a59e86'
  error: '#ba1a1a'
  on-error: '#ffffff'
  error-container: '#ffdad6'
  on-error-container: '#93000a'
  primary-fixed: '#cae8e8'
  primary-fixed-dim: '#aecccc'
  on-primary-fixed: '#022020'
  on-primary-fixed-variant: '#304c4b'
  secondary-fixed: '#ffdbc7'
  secondary-fixed-dim: '#ecbda0'
  on-secondary-fixed: '#2e1504'
  on-secondary-fixed-variant: '#603f29'
  tertiary-fixed: '#ebe2c8'
  tertiary-fixed-dim: '#cec6ad'
  on-tertiary-fixed: '#1f1c0b'
  on-tertiary-fixed-variant: '#4c4733'
  background: '#fbf9f5'
  on-background: '#1b1c1a'
  surface-variant: '#e4e2de'
typography:
  headline-lg:
    fontFamily: Plus Jakarta Sans
    fontSize: 40px
    fontWeight: '700'
    lineHeight: '1.2'
    letterSpacing: -0.02em
  headline-lg-mobile:
    fontFamily: Plus Jakarta Sans
    fontSize: 32px
    fontWeight: '700'
    lineHeight: '1.2'
    letterSpacing: -0.02em
  headline-md:
    fontFamily: Plus Jakarta Sans
    fontSize: 28px
    fontWeight: '600'
    lineHeight: '1.3'
    letterSpacing: -0.01em
  body-lg:
    fontFamily: Plus Jakarta Sans
    fontSize: 18px
    fontWeight: '400'
    lineHeight: '1.6'
    letterSpacing: 0.01em
  body-md:
    fontFamily: Plus Jakarta Sans
    fontSize: 16px
    fontWeight: '400'
    lineHeight: '1.6'
    letterSpacing: 0.01em
  label-md:
    fontFamily: Plus Jakarta Sans
    fontSize: 14px
    fontWeight: '600'
    lineHeight: '1.4'
    letterSpacing: 0.05em
  label-sm:
    fontFamily: Plus Jakarta Sans
    fontSize: 12px
    fontWeight: '500'
    lineHeight: '1.4'
    letterSpacing: 0.03em
rounded:
  sm: 0.5rem
  DEFAULT: 1rem
  md: 1.5rem
  lg: 2rem
  xl: 3rem
  full: 9999px
spacing:
  base: 8px
  container-padding-mobile: 24px
  container-padding-desktop: 64px
  stack-gap: 16px
  section-gap: 48px
---

## Brand & Style

This design system is built upon the principle of *Rifqah*—companionship characterized by gentleness. It serves as a digital sanctuary for Arab and Egyptian communities, prioritizing psychological safety and absolute anonymity. The visual language moves away from the clinical, cold aesthetic of traditional security apps, opting instead for a "Warm Organicism" that feels like a private, sun-drenched courtyard.

The style avoids the "move fast and break things" energy of typical tech, favoring a slow, intentional, and protective interface. High-security features are tucked behind a serene facade of soft textures and generous negative space, ensuring that users feel held rather than monitored.

## Colors

The palette is derived from the natural landscapes of the Middle East—deep pine forests, desert sands, and sun-bleached limestone.

- **Background (#FDFBF7):** A warm, off-white "Cream" that eliminates eye strain and feels more human than stark digital white.
- **Primary (#1E3A3A):** A "Deep Sage Pine" used for all primary typography and structural elements. It conveys depth, stability, and institutional-grade security without the coldness of black or navy.
- **Secondary (#D9AB8F):** "Terracotta Sand" is used for call-to-actions and interactive highlights. It provides a soft, warm contrast that feels inviting.
- **Tertiary (#F4EBD0):** A muted "Oatmeal" for secondary surfaces and containers to create depth without using gray.

## Typography

This design system utilizes **Plus Jakarta Sans** for its balanced, modern, and friendly geometry. To emphasize the premium and serene nature of the platform, we employ wide line-heights (1.6 for body) and generous letter spacing for labels.

Text should never be pure black; use the Primary Deep Sage Pine to maintain a soft, organic feel. Headlines are set with slight negative letter spacing to feel "locked in" and secure, while body text is spaced out to promote legibility and a sense of calm during difficult conversations.

## Layout & Spacing

The layout philosophy is based on **expansive sanctuary.** We prioritize "breathable" interfaces over information density.

- **Grid:** A 12-column fluid grid for desktop and a 4-column grid for mobile.
- **Margins:** Unusually large side margins (24px on mobile) to create a centered, focused experience that feels private.
- **Gaps:** Use the `section-gap` (48px) liberally to separate distinct phases of user interaction (e.g., identity verification vs. voice room entry).
- **Rhythm:** Elements are spaced using an 8px baseline, with a preference for larger increments (24px, 32px, 48px) to reinforce the serene aesthetic.

## Elevation & Depth

Elevation is achieved through **Tonal Layering** rather than traditional drop shadows. By stacking slightly darker or warmer shades of the background (Tertiary Oatmeal #F4EBD0), we create a sense of physical objects resting on a surface.

Where shadows are necessary for interactivity (like floating action buttons), use **Ambient Bloom Shadows**:
- Very large blur radii (32px+).
- Very low opacity (8-12%).
- Tinted with the Primary color (#1E3A3A) to ensure the shadow feels like a natural part of the environment, never a "dirty" gray.
- **Backdrop Blurs:** Use subtle frosted glass effects on navigation overlays to maintain awareness of the "safe space" behind the current action.

## Shapes

The shape language is **hyper-organic.** There are no sharp corners in this design system. All containers, buttons, and input fields utilize pill-shaped or deeply rounded corners (minimum 24px).

This roundness mimics the softness of human forms and natural elements like river stones, psychologically reinforcing the brand values of gentleness and companionship. Even "square" elements like cards should have a minimum radius of `rounded-xl` (48px) to maintain the protective, soft-shell feel.

## Components

### Buttons
- **Primary:** Deep Sage Pine background with Cream text. Pill-shaped. No border.
- **Secondary:** Terracotta Sand background with Deep Sage Pine text. Use for encouraging actions (e.g., "Join Conversation").
- **Ghost:** Transparent background with Primary color border (2px) and text.

### Cards & Containers
- Cards should use a subtle Tertiary color background (#F4EBD0) to distinguish them from the main page surface.
- Corners are always `rounded-xl`. 
- Content inside cards should have a minimum of 32px internal padding.

### Input Fields
- Inputs are designed as soft troughs. Background should be slightly darker than the page (#F4EBD0).
- Borders only appear on focus, using the Terracotta Sand color to indicate "warmth" and activity.
- Placeholder text should be the Primary color at 50% opacity.

### Voice & Privacy Indicators
- **The "Safe Pulse":** A slow, breathing animation (opacity 1.0 to 0.4) using the Primary color to indicate an active, encrypted voice connection.
- **Anonymity Shield:** A small, permanent organic icon at the top of every screen to remind users that their identity is shielded.

### Lists
- Items are separated by generous vertical spacing (16px) rather than divider lines.
- Each list item is its own rounded container to avoid a "data table" feel.
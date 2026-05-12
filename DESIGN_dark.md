---
name: Strategic Enterprise Identity (Dark)
colors:
  surface: '#121414'
  surface-dim: '#121414'
  surface-bright: '#38393a'
  surface-container-lowest: '#0c0f0f'
  surface-container-low: '#1a1c1c'
  surface-container: '#1e2020'
  surface-container-high: '#282a2b'
  surface-container-highest: '#333535'
  on-surface: '#e2e2e2'
  on-surface-variant: '#c5c6cd'
  inverse-surface: '#e2e2e2'
  inverse-on-surface: '#2f3131'
  outline: '#8f9097'
  outline-variant: '#44474d'
  surface-tint: '#b9c7e4'
  primary: '#b9c7e4'
  on-primary: '#233148'
  primary-container: '#0a192f'
  on-primary-container: '#74829d'
  inverse-primary: '#515f78'
  secondary: '#adc7ff'
  on-secondary: '#002e68'
  secondary-container: '#4a8eff'
  on-secondary-container: '#00285b'
  tertiary: '#a0cbf3'
  on-tertiary: '#003351'
  tertiary-container: '#001b2d'
  on-tertiary-container: '#5b86ab'
  error: '#ffb4ab'
  on-error: '#690005'
  error-container: '#93000a'
  on-error-container: '#ffdad6'
  primary-fixed: '#d6e3ff'
  primary-fixed-dim: '#b9c7e4'
  on-primary-fixed: '#0d1c32'
  on-primary-fixed-variant: '#39475f'
  secondary-fixed: '#d8e2ff'
  secondary-fixed-dim: '#adc7ff'
  on-secondary-fixed: '#001a41'
  on-secondary-fixed-variant: '#004493'
  tertiary-fixed: '#cce5ff'
  tertiary-fixed-dim: '#a0cbf3'
  on-tertiary-fixed: '#001d31'
  on-tertiary-fixed-variant: '#1a4a6c'
  background: '#121414'
  on-background: '#e2e2e2'
  surface-variant: '#333535'
typography:
  display-lg:
    fontFamily: Plus Jakarta Sans
    fontSize: 48px
    fontWeight: '700'
    lineHeight: '1.1'
    letterSpacing: -0.02em
  display-lg-mobile:
    fontFamily: Plus Jakarta Sans
    fontSize: 36px
    fontWeight: '700'
    lineHeight: '1.2'
  headline-lg:
    fontFamily: Plus Jakarta Sans
    fontSize: 32px
    fontWeight: '600'
    lineHeight: '1.2'
  headline-md:
    fontFamily: Plus Jakarta Sans
    fontSize: 24px
    fontWeight: '600'
    lineHeight: '1.3'
  body-lg:
    fontFamily: Plus Jakarta Sans
    fontSize: 18px
    fontWeight: '400'
    lineHeight: '1.6'
  body-md:
    fontFamily: Plus Jakarta Sans
    fontSize: 16px
    fontWeight: '400'
    lineHeight: '1.6'
  body-sm:
    fontFamily: Plus Jakarta Sans
    fontSize: 14px
    fontWeight: '400'
    lineHeight: '1.5'
  label-lg:
    fontFamily: Plus Jakarta Sans
    fontSize: 14px
    fontWeight: '600'
    lineHeight: '1'
    letterSpacing: 0.05em
  label-sm:
    fontFamily: Plus Jakarta Sans
    fontSize: 12px
    fontWeight: '500'
    lineHeight: '1'
rounded:
  sm: 0.25rem
  DEFAULT: 0.5rem
  md: 0.75rem
  lg: 1rem
  xl: 1.5rem
  full: 9999px
spacing:
  base: 8px
  xs: 4px
  sm: 12px
  md: 24px
  lg: 48px
  xl: 80px
  gutter: 24px
  margin-mobile: 16px
  margin-desktop: 64px
---

## Brand & Style

The design system is engineered for a high-stakes consultancy environment, blending the authority of established enterprise tech with the agility of modern innovation. This version utilizes a **Premium Dark** aesthetic, evoking a sense of stability through a dominant Deep Navy foundation, while signaling technological forward-momentum with Electric Blue accents.

The chosen style is **Corporate / Modern (Dark Mode)**, characterized by exceptional clarity, intentional use of "darkspace," and a high-contrast environment. To differentiate from traditional consultancy aesthetics, this design system introduces **subtle Glassmorphism** in overlay elements and **luminous accents** to create depth and a "digital-first" feel. The overall response should be one of competence, precision, and architectural integrity.

## Colors

The color palette is architected for maximum legibility and professional hierarchy in a dark-mode environment. 

- **Primary (Deep Navy):** #0A192F. Used as the system's foundational canvas and primary background color. It provides the "anchor" for the entire system.
- **Secondary (Electric Blue):** #007BFF. Reserved for high-priority actions, interactive states, and data visualization highlights. It acts as the "innovative spark" against the dark tones.
- **Tertiary (Medium Navy):** #003B5C. Utilized for structural accents, secondary buttons, or surface-level containers to create hierarchical depth.
- **Neutral (Cool Gray/White):** The system relies on Light Gray (#F4F4F4) for typography and iconography to ensure high contrast against the dark background.

Status colors (Success, Warning, Error) are adjusted for vibrancy and high contrast against the Deep Navy background.

## Typography

This design system utilizes **Plus Jakarta Sans** across all levels to ensure a cohesive, modern, and highly readable experience. The type scale is designed with a strong vertical rhythm, using a 1.6x line height for body copy.

- **Headlines:** Use tighter letter spacing and heavier weights (600-700) to create a commanding presence against the dark background.
- **Body:** Standardized at 16px. In dark mode, light text on dark backgrounds can appear "thicker," so weights are carefully managed for clarity.
- **Labels:** Small caps or increased letter spacing should be applied to labels to differentiate them from functional body text.
- **Mobile scaling:** Display sizes are reduced by approximately 25% on mobile devices to maintain screen real estate while preserving hierarchy.

## Layout & Spacing

The layout philosophy follows a **Fixed Grid** approach for desktop views to maintain a curated, editorial feel suitable for high-end consultancy, transitioning to a **Fluid Grid** for mobile devices.

- **Grid System:** A 12-column grid is used for desktop (max-width: 1440px) with 24px gutters.
- **Spacing Rhythm:** Based on an 8px linear scale. All margins and paddings must be multiples of 8 to ensure mathematical harmony.
- **Breakpoints:**
  - **Mobile:** < 600px (4 columns)
  - **Tablet:** 600px - 1024px (8 columns)
  - **Desktop:** > 1024px (12 columns)
- **Safe Areas:** Generous 64px external margins on desktop prevent content from feeling "crowded," reinforcing the premium nature of the brand.

## Elevation & Depth

Visual hierarchy in this dark design system is achieved through **Tonal Layering** and **Subtle Inner Glows**. Instead of shadows, surfaces are distinguished by their brightness relative to the base:

- **Level 0 (Base):** #0A192F background.
- **Level 1 (Cards/Containers):** Medium Navy (#003B5C) surfaces with a very subtle, low-opacity Neutral border (1px).
- **Level 2 (Dropdowns/Modals):** Lighter navy tones with high-diffusion, Electric Blue tinted shadows (Blur: 32px, Y: 12px, Color: rgba(0, 123, 255, 0.12)).
- **Glassmorphism:** Navigation bars and floating action panels use a backdrop-blur (12px) with a semi-transparent navy fill (opacity: 70%) to maintain context.

## Shapes

The shape language is **Rounded (Level 2)**. This specific radius (8px/0.5rem) was chosen to soften the "industrial" feel of enterprise tech while remaining structured enough for professional consultancy.

- **Components:** Buttons, input fields, and cards utilize the base 8px radius.
- **Large Elements:** Larger containers can scale up to `rounded-xl` (24px) for a more modern feel.
- **Iconography:** Icons should feature slightly rounded terminals and a consistent 2px stroke weight to match the typeface's geometry.

## Components

### Buttons
- **Primary:** Electric Blue (#007BFF) fill for high visibility. White text.
- **Secondary:** Medium Navy (#003B5C) background with a Neutral (#F4F4F4) outline (1px).
- **Ghost:** No fill, Neutral text, used for low-emphasis actions.

### Inputs
- **Text Fields:** Deep Navy background, 1px border (#ffffff20). On focus, the border transitions to Electric Blue (#007BFF) with a 2px outer glow.
- **Labels:** Always positioned above the field in `label-sm` style, using the Neutral color at 80% opacity.

### Cards
- **Enterprise Cards:** Medium Navy (#003B5C) background, 8px corner radius.
- **Feature Cards:** Feature a 4px top-border accent in Electric Blue to draw attention to "Innovation" metrics.

### Lists & Data
- **Data Tables:** Clean, no vertical lines. Horizontal dividers in #ffffff10. Header rows utilize a semi-transparent Electric Blue background with white text for structural clarity.
- **Chips:** Medium Navy (#003B5C) backgrounds with Neutral text for tags; Electric Blue backgrounds for "Active" status indicators.
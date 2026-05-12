---
name: Strategic Enterprise Identity
colors:
  surface: '#f9f9f9'
  surface-dim: '#dadada'
  surface-bright: '#f9f9f9'
  surface-container-lowest: '#ffffff'
  surface-container-low: '#f3f3f3'
  surface-container: '#eeeeee'
  surface-container-high: '#e8e8e8'
  surface-container-highest: '#e2e2e2'
  on-surface: '#1a1c1c'
  on-surface-variant: '#44474d'
  inverse-surface: '#2f3131'
  inverse-on-surface: '#f1f1f1'
  outline: '#75777e'
  outline-variant: '#c5c6cd'
  surface-tint: '#515f78'
  primary: '#000000'
  on-primary: '#ffffff'
  primary-container: '#0d1c32'
  on-primary-container: '#76849f'
  inverse-primary: '#b9c7e4'
  secondary: '#0059bb'
  on-secondary: '#ffffff'
  secondary-container: '#0070ea'
  on-secondary-container: '#fefcff'
  tertiary: '#000000'
  on-tertiary: '#ffffff'
  tertiary-container: '#001d31'
  on-tertiary-container: '#5d88ad'
  error: '#ba1a1a'
  on-error: '#ffffff'
  error-container: '#ffdad6'
  on-error-container: '#93000a'
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
  background: '#f9f9f9'
  on-background: '#1a1c1c'
  surface-variant: '#e2e2e2'
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

The design system is engineered for a high-stakes consultancy environment, blending the authority of established enterprise tech with the agility of modern innovation. The visual language centers on **Trustworthy Professionalism**, evoking a sense of stability through a dominant Deep Navy foundation, while signaling technological forward-momentum with Electric Blue accents.

The chosen style is **Corporate / Modern**, characterized by exceptional clarity, intentional whitespace, and a high-contrast environment. To differentiate from traditional, static consultancy aesthetics, this design system introduces **subtle Glassmorphism** in overlay elements and **refined gradients** to create depth and a "digital-first" feel. The overall response should be one of competence, precision, and architectural integrity.

## Colors

The color palette is architected for maximum legibility and professional hierarchy. 

- **Primary (Deep Navy):** Used for core branding, navigation backgrounds, and primary headings. It provides the "anchor" for the entire system.
- **Secondary (Electric Blue):** Reserved for high-priority actions, interactive states, and data visualization highlights. It acts as the "innovative spark" against the darker tones.
- **Tertiary (Medium Navy):** Utilized for structural accents, secondary buttons, or illustrative icons to bridge the gap between the deep primary and light neutrals.
- **Neutral (Light Gray/White):** The canvas is predominantly white (#FFFFFF) with Light Gray (#F4F4F4) used for section backgrounds and container fills to define spatial boundaries without adding visual noise.

Status colors (Success, Warning, Error) should follow standard industry conventions but be adjusted for high contrast against the light background.

## Typography

This design system utilizes **Plus Jakarta Sans** across all levels to ensure a cohesive, modern, and highly readable experience. The type scale is designed with a strong vertical rhythm, using a 1.6x line height for body copy to ensure readability in data-dense consultancy reports.

- **Headlines:** Use tighter letter spacing and heavier weights (600-700) to create a commanding presence.
- **Body:** Standardized at 16px for optimal balance between information density and clarity.
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

Visual hierarchy in this design system is achieved through **Tonal Layering** and **Ambient Shadows**. Instead of heavy borders, surfaces are distinguished by their elevation levels:

- **Level 0 (Base):** #FFFFFF or #F4F4F4 background.
- **Level 1 (Cards/Containers):** Pure white surfaces with an extremely soft, diffused shadow (Blur: 16px, Y: 4px, Color: rgba(10, 25, 47, 0.06)).
- **Level 2 (Dropdowns/Modals):** High-diffusion shadows with a subtle navy tint (Blur: 32px, Y: 12px, Color: rgba(10, 25, 47, 0.12)).
- **Glassmorphism:** Navigation bars and floating action panels should use a backdrop-blur (12px) with a semi-transparent white fill (opacity: 80%) to maintain context of the underlying content.

## Shapes

The shape language is **Rounded (Level 2)**. This specific radius (8px/0.5rem) was chosen to soften the "industrial" feel of enterprise tech while remaining structured enough for professional consultancy.

- **Components:** Buttons, input fields, and cards utilize the base 8px radius.
- **Large Elements:** Larger containers like hero images or marketing cards can scale up to `rounded-xl` (24px) for a more modern, approachable feel.
- **Iconography:** Icons should feature slightly rounded terminals and a consistent 2px stroke weight to match the typeface's geometry.

## Components

### Buttons
- **Primary:** Deep Navy (#0A192F) fill with a subtle vertical gradient to Medium Navy (#003B5C). White text.
- **Secondary:** Electric Blue (#007BFF) outline (2px) with matching text.
- **Ghost:** No fill, Primary Navy text, used for low-emphasis actions.

### Inputs
- **Text Fields:** White background, 1px border (#D1D5DB). On focus, the border transitions to Electric Blue (#007BFF) with a 2px outer glow.
- **Labels:** Always positioned above the field in `label-sm` style.

### Cards
- **Enterprise Cards:** White background, 8px corner radius, and Level 1 ambient shadow. Used for services, case studies, and team profiles.
- **Feature Cards:** Feature a 4px top-border accent in Electric Blue to draw attention to "Innovation" metrics.

### Lists & Data
- **Data Tables:** Clean, no vertical lines. Horizontal dividers in #F4F4F4. Header rows utilize the Deep Navy background with white text for maximum structural clarity.
- **Chips:** Soft-gray (#F4F4F4) backgrounds with Primary Navy text for tags; Electric Blue backgrounds for "Active" status indicators.

### Iconography
- Use crisp, dual-tone or line-based icons. Key metaphors should incorporate Electric Blue accents to highlight the "point of innovation" within the graphic.
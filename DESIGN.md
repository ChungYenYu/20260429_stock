---
version: alpha
name: Apple Premium
description: Cinematic minimalism, high contrast, and the signature Apple Blue.
colors:
  primary: "#000000"
  secondary: "#f5f5f7"
  accent: "#0071e3"
  text-primary: "#1d1d1f"
  text-secondary: "rgba(0, 0, 0, 0.8)"
  link: "#0066cc"
  link-dark: "#2997ff"
typography:
  display-hero:
    fontFamily: "system-ui, -apple-system"
    fontSize: "3.5rem"
    fontWeight: 600
    lineHeight: 1.07
    letterSpacing: "-0.02em"
  section-heading:
    fontFamily: "system-ui, -apple-system"
    fontSize: "2.5rem"
    fontWeight: 600
    lineHeight: 1.1
  body-md:
    fontFamily: "system-ui, -apple-system"
    fontSize: "1.06rem"
    lineHeight: 1.47
    letterSpacing: "-0.022em"
rounded:
  sm: "5px"
  md: "8px"
  lg: "12px"
  pill: "980px"
spacing:
  base: "8px"
components:
  button-primary:
    backgroundColor: "{colors.accent}"
    textColor: "#FFFFFF"
    rounded: "{rounded.md}"
    padding: "8px 15px"
  navbar:
    backgroundColor: "rgba(0, 0, 0, 0.8)"
    blur: "20px"
    height: "48px"
---

## Overview
This design system follows Apple's "Product-as-Hero" philosophy.

## Colors
- **Apple Blue (#0071e3):** The only chromatic color, used exclusively for interaction.
- **Light Gray (#f5f5f7):** Soft neutral for sections.

## Typography
Uses system fonts (SF Pro equivalents) with tight tracking and high contrast line-heights.

## Components
- **Pill CTA:** Use `rounded.pill` for "Learn More" links.
- **Glass Nav:** Translucent background with backdrop blur.

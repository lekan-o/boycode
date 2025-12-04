# Frontend Mentor - Blog preview card solution

This is a solution to the [Blog preview card challenge on Frontend Mentor](https://www.frontendmentor.io/challenges/blog-preview-card-ckPaj01IcS). Frontend Mentor challenges help you improve your coding skills by building realistic projects.

## Table of contents

- [Overview](#overview)
  - [The challenge](#the-challenge)
  - [Screenshot](#screenshot)
  - [Links](#links)
- [My process](#my-process)
  - [Built with](#built-with)
  - [What I learned](#what-i-learned)
  - [Continued development](#continued-development)
  - [Useful resources](#useful-resources)
- [Author](#author)

## Overview

### The challenge

Users should be able to:

- View the optimal layout for the interface depending on their device's screen size
- See hover and focus states for all interactive elements on the page

### Screenshot

![Blog preview card - Desktop view](./preview.jpg)

*A clean, modern blog preview card with hover effects and responsive design*

### Links

- Solution URL: [GitHub Repository](https://github.com/lekan-o/boycode)
- Live Site URL: [View Live Demo](https://lekan-o.github.io/boycode)

## My process

### Built with

- Semantic HTML5 markup
- CSS custom properties for colors and typography
- Flexbox for layout
- Mobile-first workflow
- CSS transitions and hover effects
- Google Fonts (Figtree)
- WCAG accessibility standards

### What I learned

This project reinforced my understanding of several key web development concepts:

**CSS Custom Properties for Design Consistency**
Using CSS variables made it easy to maintain the design system across the project:

```css
body {
  --color-yellow: hsl(47, 88%, 63%);
  --color-white: hsl(0, 0%, 100%);
  --color-dark: hsl(0, 0%, 7%);
  --color-gray: hsl(0, 0%, 42%);
}
```

**Interactive Element Styling**
Implementing smooth transitions and hover states improved user experience:

```css
.card:hover {
  box-shadow: 0 8px 0 0 hsl(0, 0%, 7%);
  transform: translateY(-4px);
  transition: box-shadow 0.2s ease, transform 0.2s ease;
}
```

**Accessibility Focus States**
Ensuring all interactive elements have clear focus states for keyboard navigation:

```css
.article-title:focus {
  outline: 2px solid hsl(0, 0%, 7%);
  outline-offset: 2px;
}
```

This project demonstrated the importance of paying attention to the small details that make a polished user experience, from box shadows to transition timing to focus indicators.

### Continued development

In future projects, I'd like to:

- Explore CSS animations for more complex interactive states
- Implement progressive enhancement techniques
- Practice building component libraries with reusable styles
- Deepen my understanding of CSS Grid for more complex layouts
- Implement more advanced accessibility features like ARIA labels

### Useful resources

- [MDN Web Docs - CSS Transitions](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Transitions) - Excellent reference for understanding CSS transitions and timing functions
- [Web Content Accessibility Guidelines (WCAG)](https://www.w3.org/WAI/WCAG21/quickref/) - Helpful for understanding focus states and keyboard navigation requirements
- [Google Fonts - Figtree](https://fonts.google.com/specimen/Figtree) - The typeface used in this design
- [Frontend Mentor](https://www.frontendmentor.io) - Great platform for building real-world projects

## Author

- Frontend Mentor - [@lekan-o](https://www.frontendmentor.io/profile/devv-leo)
- GitHub - [lekan-o](https://github.com/lekan-o)

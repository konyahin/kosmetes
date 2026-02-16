(function() {
    'use strict';

    const starsContainer = document.createElement('div');
    starsContainer.id = 'stars';
    starsContainer.style.cssText = `
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        pointer-events: none;
        z-index: 0;
        overflow: hidden;
    `;
    document.body.appendChild(starsContainer);

    function createStar() {
        const star = document.createElement('div');
        const size = Math.random() * 2 + 1;
        const x = Math.random() * 100;
        const y = Math.random() * 100;
        const duration = Math.random() * 3 + 2;
        const delay = Math.random() * 2;
        const opacity = Math.random() * 0.7 + 0.3;

        star.style.cssText = `
            position: absolute;
            width: ${size}px;
            height: ${size}px;
            background-color: white;
            border-radius: 50%;
            left: ${x}%;
            top: ${y}%;
            opacity: ${opacity};
            animation: twinkle ${duration}s ease-in-out infinite;
            animation-delay: ${delay}s;
        `;

        return star;
    }

    const style = document.createElement('style');
    style.textContent = `
        @keyframes twinkle {
            0%, 100% { opacity: 0.3; }
            50% { opacity: 1; }
        }
    `;
    document.head.appendChild(style);

    const starCount = 75;
    for (let i = 0; i < starCount; i++) {
        starsContainer.appendChild(createStar());
    }

    function isDarkTheme() {
        return window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;
    }

    function updateStarsVisibility() {
        starsContainer.style.display = isDarkTheme() ? 'block' : 'none';
    }

    if (window.matchMedia) {
        window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', updateStarsVisibility);
    }

    updateStarsVisibility();
})();

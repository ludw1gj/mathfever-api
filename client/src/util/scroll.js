/**
 * Scroll to top of element in a linear fashion.
 * @param {HTMLElement} element - The element to scroll.
 * @param {number} scrollDuration - The scroll duration in milliseconds.
 */
export const scrollToTop = (element, scrollDuration) => {
  const scrollStep = -element.scrollTop / (scrollDuration / 15);

  let scrollInterval = setInterval(() => {
    if (element.scrollTop !== 0) {
      element.scrollBy(0, scrollStep);
    } else clearInterval(scrollInterval);
  }, 15);
};

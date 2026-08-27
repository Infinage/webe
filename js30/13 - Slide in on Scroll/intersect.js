document.addEventListener("alpine:init", () => {
  Alpine.data("visible_on_scrolled", () => ({
    isVisible: false,
    init() {
      const obs = new IntersectionObserver(([entry]) => {
        this.isVisible = entry.isIntersecting;
        console.log(this.isVisible);
      }, { threshold: .15 });
      obs.observe(this.$el);
    }
  }));
});

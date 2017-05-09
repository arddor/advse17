<template>
  <div id="app">
    <section class="hero is-primary">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">Twitter Popularity</h1>
          <h2 class="subtitle">Track the popularity of brands on twitter</h2>
        </div>
      </div>
    </section>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'app'
}
</script>

<style lang="scss">
  @import "./variables";
  @import '~bulma';
</style>

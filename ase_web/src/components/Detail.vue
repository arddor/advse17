<template>
  <div class="section" v-if="term">
    <h1 class="title">{{term.term}}</h1>
  </div>
  <div class="section" v-else>
    <p>Ups something went wrong</p>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  async beforeRouteEnter (to, from, next) {
    try {
      let {data} = await axios.get(`api/terms/${to.params.id}`)
      next(vm => {
        vm.$data.term = data
      })
    } catch (e) {
      next()
    }
  },
  data () {
    return {
      term: null
    }
  }
}
</script>

<style lang="css">
</style>

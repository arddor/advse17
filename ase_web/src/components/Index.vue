<template>
  <section class="hero is-medium" v-if="loading">
    <div class="hero-body">
      <div class="container">
        <div class="column is-one-third is-offset-one-third has-text-centered">
          <pulse-loader :loading="loading"></pulse-loader>
          <h1 class="title">loading</h1>
        </div>
      </div>
    </div>
  </section>
  <section class="section" v-else>
    <div class="field has-addons">
      <p class="control is-expanded">
        <input class="input" type="text" placeholder="Term to track" v-model="term">
      </p>
      <p class="control">
        <a class="button is-info" @click="createTerm">Add Term</a>
      </p>
    </div>

    <ul v-for="term in terms">
      <li>
        <a class="box" @click="goToDetails(term.id)">
          {{term.term}}
        </a>
      </li>
    </ul>
  </section>
</template>

<script>
import axios from 'axios'
import PulseLoader from 'vue-spinner/src/PulseLoader.vue'
export default {
  name: 'Index',
  components: {
    PulseLoader
  },
  data () {
    return {
      loading: false,
      term: '',
      terms: []
    }
  },
  async created () {
    this.loading = true
    try {
      let {data} = await axios.get('api/terms')
      this.terms = data
    } catch (e) {
      // TODO: Show popup error message
      console.log(e)
    }
    this.loading = false
  },
  methods: {
    goToDetails (id) {
      this.$router.push({name: 'Detail', params: { id }})
    },
    async createTerm () {
      if (this.term !== '') {
        try {
          let {data} = await axios.post('api/terms', {term: this.term})
          this.terms.push(data)
          this.term = ''
        } catch (e) {
          console.log(e)
          // TODO: Show popup error message
        }
      }
      // TODO: Show popup that term should not be empty
    }
  }
}
</script>

<style lang="scss" scoped>
  li{
    margin: 10px 0;
  }
</style>

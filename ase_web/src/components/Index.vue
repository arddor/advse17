<template>
  <section class="section">
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
export default {
  name: 'Index',
  data () {
    return {
      term: '',
      terms: []
    }
  },
  async beforeRouteEnter (to, from, next) {
    try {
      let {data} = await axios.get('api/terms')
      next(vm => {
        vm.$data.terms = data
      })
    } catch (e) {
      next()
    }
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

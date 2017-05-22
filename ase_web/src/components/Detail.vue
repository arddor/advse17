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
    <h1 class="title" v-if="term">{{term}}</h1>
    <plotly
      v-if="chart"
      :data="chart"
      :layout="{barmode: 'relative'}"
    ></plotly>
  </section>
</template>

<script>
import axios from 'axios'
import Plotly from '@/components/Plotly'
import PulseLoader from 'vue-spinner/src/PulseLoader.vue'
export default {
  components: {
    Plotly,
    PulseLoader
  },
  async created () {
    let {id} = this.$route.params
    this.loading = true
    try {
      let {data} = await axios.get(`api/terms/${id}`)
      this.term = data.term
      this.chart = this.chartData(data.data)
      // this.chartData(data.data)
    } catch (e) {
      console.log(e)
    }
    this.loading = false

    var ws = new WebSocket('ws://127.0.0.1:5002/echo')
    ws.onopen = evt => console.log(evt)
    ws.onclose = evt => console.log('closed')
    ws.onmessage = evt => console.log(evt)
  },
  data () {
    return {
      loading: false,
      term: null,
      chart: null
    }
  },
  methods: {
    chartData (data) {
      let love = {x: [], y: [], type: 'bar', name: 'love'}
      let positive = {x: [], y: [], type: 'bar', name: 'positive'}
      let negative = {x: [], y: [], type: 'bar', name: 'negative'}
      let hate = {x: [], y: [], type: 'bar', name: 'hate'}

      data.forEach(item => {
        if (item.sentiment > 0.75) {
          love.x.push(new Date(item.time))
          love.y.push(item.sentiment)
        } else if (item.sentiment > 0.5) {
          positive.x.push(new Date(item.time))
          positive.y.push(item.sentiment)
        } else if (item.sentiment > 0.25) {
          negative.x.push(new Date(item.time))
          negative.y.push(item.sentiment - 1)
        } else if (item.sentiment < 0.25) {
          hate.x.push(new Date(item.time))
          hate.y.push(item.sentiment - 1)
        }
      })
      return [love, positive, negative, hate]
    },
    testChart (data) {
      let x = []
      let y = []

      data.forEach(item => {
        x.push(new Date(item.time))
        y.push(item.sentiment)
      })

      return [{x, y, type: 'bar'}]
    }
  }
}
</script>

<style lang="css">
</style>

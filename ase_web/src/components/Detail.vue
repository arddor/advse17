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
    <nav class="level">
      <div class="level-item has-text-centered">
        <div>
          <p class="heading">Period</p>
          <p class="title">
            <div class="field">
              <p class="control">
                <span class="select">
                  <select v-model.number="selected">
                    <option value="3600">1h</option>
                    <option value="43200">12h</option>
                    <option value="86400">1d</option>
                    <option value="604800">1 week</option>
                  </select>
                </span>
              </p>
            </div>
          </p>
        </div>
      </div>
      <div class="level-item has-text-centered">
        <div>
          <p class="heading">Average</p>
          <p class="title">{{average}}</p>
        </div>
      </div>
    </nav>
    <plotly
      v-if="chart"
      :data="chart"
      :layout="{barmode: 'relative', margin: {l: 50, r: 50, b: 50, t: 50,pad: 4}}"
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
    this.loadTerm()

    var ws = new WebSocket('ws://127.0.0.1:5002/echo')
    ws.onopen = evt => console.log(evt)
    ws.onclose = evt => console.log('closed')
    ws.onmessage = evt => console.log(evt)
  },
  data () {
    return {
      loading: false,
      selected: 3600, // seconds to show
      average: 0,
      term: null,
      chart: null
    }
  },
  watch: {
    selected (val) {
      this.loadTerm(val)
    }
  },
  methods: {
    chartData (data) {
      let love = {x: [], y: [], type: 'bar', name: 'love', marker: {color: 'rgb(12, 134, 35)'}}
      let positive = {x: [], y: [], type: 'bar', name: 'positive', marker: {color: 'rgb(38, 173, 203)'}}
      let negative = {x: [], y: [], type: 'bar', name: 'negative', marker: {color: 'rgb(245, 139, 28)'}}
      let hate = {x: [], y: [], type: 'bar', name: 'hate', marker: {color: 'rgb(232, 20, 20)'}}

      data.forEach(item => {
        if (item.sentiment > 0.75) {
          love.x.push(new Date(item.time))
          love.y.push(1)
        } else if (item.sentiment > 0.5) {
          positive.x.push(new Date(item.time))
          positive.y.push(0.5)
        } else if (item.sentiment > 0.25) {
          negative.x.push(new Date(item.time))
          negative.y.push(-0.5)
        } else if (item.sentiment < 0.25) {
          hate.x.push(new Date(item.time))
          hate.y.push(-1)
        }
      })
      return [love, positive, negative, hate]
    },
    async loadTerm (seconds = 3600) {
      let {id} = this.$route.params
      this.loading = true
      try {
        let {data} = await axios.get(`api/terms/${id}`, {params: {seconds}})
        this.term = data.term
        this.length = data.data.length
        this.chart = this.chartData(data.data)
        this.average = data.data.length === 0
        ? 0
        : data.data.reduce((total, num) => ({sentiment: total.sentiment + num.sentiment})).sentiment / data.data.length
        // this.chartData(data.data)
      } catch (e) {
        console.log(e)
      }
      this.loading = false
    }
  }
}
</script>

<style lang="css">
</style>

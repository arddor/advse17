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
  created () {
    let {id} = this.$route.params
    this.loadTerm()

    let {port} = window.location
    let {host} = window.location
    this.ws = new window.WebSocket(`ws://${host}:${port}/ws/changes/${id}`)
    this.ws.onopen = evt => console.log('opened websocket')
    this.ws.onclose = evt => console.log('closed')
    this.ws.onmessage = evt => {
      let json = JSON.parse(evt.data)
      if (json.id === id) {
        this.sum += json.data.sentiment
        this.length += 1
        if (json.data.sentiment > 0.75) {
          this.chart[0].x.push(new Date(json.data.timestamp))
          this.chart[0].y.push(1)
        } else if (json.data.sentiment > 0.5) {
          this.chart[1].x.push(new Date(json.data.timestamp))
          this.chart[1].y.push(0.5)
        } else if (json.data.sentiment > 0.25) {
          this.chart[2].x.push(new Date(json.data.timestamp))
          this.chart[2].y.push(-0.5)
        } else if (json.data.sentiment < 0.25) {
          this.chart[3].x.push(new Date(json.data.timestamp))
          this.chart[3].y.push(-1)
        }
      }
    }
  },
  beforeDestroy () {
    if (this.ws) {
      this.ws.close()
    }
  },
  data () {
    return {
      loading: false,
      selected: 3600, // seconds to show
      sum: 0,
      length: 0,
      term: null,
      chart: null,
      ws: null
    }
  },
  watch: {
    selected (val) {
      this.loadTerm(val)
    }
  },
  computed: {
    average () {
      return this.length === 0
      ? 0
      : this.sum / this.length
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
        this.sum = data.data.length === 0
        ? 0
        : data.data.reduce((total, num) => ({sentiment: total.sentiment + num.sentiment})).sentiment
        // this.chartData(data.data)
      } catch (e) {
        this.$toast('Could not load data. Try to refresh the page', {
          horizontalPosition: 'center',
          verticalPosition: 'bottom',
          transition: 'slide-up'
        })
      }
      this.loading = false
    }
  }
}
</script>

<style lang="css">
</style>

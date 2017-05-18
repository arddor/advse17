<template>
  <div class="js-plotly-plot" ref="chart"></div>
</template>

<script>
import Plotly from 'plotly.js'
export default {
  props: {
    data: {
      type: Array,
      required: true,
      default: () => ([])
    },
    layout: {
      type: Object,
      default: () => ({})
    },
    config: {
      type: Object,
      default: () => ({
        modeBarButtonsToRemove: ['sendDataToCloud', 'hoverCompareCartesian'],
        displaylogo: false
      })
    },
    min: {
      type: Number,
      default: null
    },
    max: {
      type: Number,
      default: null
    },
    filename: {
      type: String,
      default: 'newplot'
    }
  },
  mounted () {
    Plotly.plot(
      this.$el,
      this.data,
      this.layout,
      this.config
    ).then(res => {
      // set filename for image download
      res.fn = this.filename
      this.chart = res
    })
    window.addEventListener('resize', this.handleResize)
  },
  beforeDestroy () {
    window.removeEventListener('resize', this.handleResize)
  },
  data () {
    return {
      chart: null
    }
  },
  computed: {
    finalLayout () {
      let layout = {
        ...this.layout,
        shapes: []
      }
      if (this.min !== null) layout.shapes.push(this.horizontalLine(this.min, '#25A9E1'))
      if (this.max !== null) layout.shapes.push(this.horizontalLine(this.max, '#FFA100'))
      return layout
    }
  },
  methods: {
    horizontalLine (y, color = 'rgb(0, 0, 0)') {
      this
      return {
        type: 'line',
        xref: 'paper',
        x0: 0,
        y0: y,
        x1: 1,
        y1: y,
        line: {color, width: 4, dash: 'dot'}
      }
    },
    handleResize (event) {
      Plotly.Plots.resize(this.chart)
    }
  },
  watch: {
    finalLayout: {
      handler (val) {
        this.$nextTick(() => {
          Plotly.relayout(this.$refs.chart, val)
        })
      },
      immediate: true
    }
  }
}
</script>

<style scoped>
.js-plotly-plot {
  max-width: 100%;
}
</style>

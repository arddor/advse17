webpackJsonp([1],{113:function(t,e){},114:function(t,e){},115:function(t,e){},116:function(t,e){},117:function(t,e){},124:function(t,e,a){a(113);var n=a(8)(a(70),a(127),null,null);t.exports=n.exports},125:function(t,e,a){a(116);var n=a(8)(a(71),a(130),"data-v-aeb8a30e",null);t.exports=n.exports},126:function(t,e,a){a(115);var n=a(8)(a(72),a(129),"data-v-70fbc797",null);t.exports=n.exports},127:function(t,e){t.exports={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return t.loading?a("section",{staticClass:"hero is-medium"},[a("div",{staticClass:"hero-body"},[a("div",{staticClass:"container"},[a("div",{staticClass:"column is-one-third is-offset-one-third has-text-centered"},[a("pulse-loader",{attrs:{loading:t.loading}}),t._v(" "),a("h1",{staticClass:"title"},[t._v("loading")])],1)])])]):a("section",{staticClass:"section"},[t.term?a("h1",{staticClass:"title"},[t._v(t._s(t.term))]):t._e(),t._v(" "),a("nav",{staticClass:"level"},[a("div",{staticClass:"level-item has-text-centered"},[a("div",[a("p",{staticClass:"heading"},[t._v("Period")]),t._v(" "),a("p",{staticClass:"title"}),a("div",{staticClass:"field"},[a("p",{staticClass:"control"},[a("span",{staticClass:"select"},[a("select",{directives:[{name:"model",rawName:"v-model.number",value:t.selected,expression:"selected",modifiers:{number:!0}}],on:{change:function(e){var a=Array.prototype.filter.call(e.target.options,function(t){return t.selected}).map(function(e){var a="_value"in e?e._value:e.value;return t._n(a)});t.selected=e.target.multiple?a:a[0]}}},[a("option",{attrs:{value:"3600"}},[t._v("1h")]),t._v(" "),a("option",{attrs:{value:"43200"}},[t._v("12h")]),t._v(" "),a("option",{attrs:{value:"86400"}},[t._v("1d")]),t._v(" "),a("option",{attrs:{value:"604800"}},[t._v("1 week")])])])])]),t._v(" "),a("p")])]),t._v(" "),a("div",{staticClass:"level-item has-text-centered"},[a("div",[a("p",{staticClass:"heading"},[t._v("Average")]),t._v(" "),a("p",{staticClass:"title"},[t._v(t._s(t.average))])])])]),t._v(" "),t.chart?a("plotly",{attrs:{data:t.chart,layout:{barmode:"relative",margin:{l:50,r:50,b:50,t:50,pad:4}}}}):t._e()],1)},staticRenderFns:[]}},128:function(t,e){t.exports={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{directives:[{name:"show",rawName:"v-show",value:t.loading,expression:"loading"}],staticClass:"v-spinner"},[a("div",{staticClass:"v-pulse v-pulse1",style:[t.spinnerStyle,t.spinnerDelay1]}),a("div",{staticClass:"v-pulse v-pulse2",style:[t.spinnerStyle,t.spinnerDelay2]}),a("div",{staticClass:"v-pulse v-pulse3",style:[t.spinnerStyle,t.spinnerDelay3]})])},staticRenderFns:[]}},129:function(t,e){t.exports={render:function(){var t=this,e=t.$createElement;return(t._self._c||e)("div",{ref:"chart",staticClass:"js-plotly-plot"})},staticRenderFns:[]}},130:function(t,e){t.exports={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return t.loading?a("section",{staticClass:"hero is-medium"},[a("div",{staticClass:"hero-body"},[a("div",{staticClass:"container"},[a("div",{staticClass:"column is-one-third is-offset-one-third has-text-centered"},[a("pulse-loader",{attrs:{loading:t.loading}}),t._v(" "),a("h1",{staticClass:"title"},[t._v("loading")])],1)])])]):a("section",{staticClass:"section"},[a("div",{staticClass:"field has-addons"},[a("p",{staticClass:"control is-expanded"},[a("input",{directives:[{name:"model",rawName:"v-model",value:t.term,expression:"term"}],staticClass:"input",attrs:{type:"text",placeholder:"Term to track"},domProps:{value:t.term},on:{input:function(e){e.target.composing||(t.term=e.target.value)}}})]),t._v(" "),a("p",{staticClass:"control"},[a("a",{staticClass:"button is-info",on:{click:t.createTerm}},[t._v("Add Term")])])]),t._v(" "),t._l(t.terms,function(e){return a("ul",[a("li",[a("a",{staticClass:"box",on:{click:function(a){t.goToDetails(e.id)}}},[t._v("\n        "+t._s(e.term)+"\n      ")])])])})],2)},staticRenderFns:[]}},131:function(t,e){t.exports={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"app"}},[t._m(0),t._v(" "),a("router-view")],1)},staticRenderFns:[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("section",{staticClass:"hero is-primary"},[a("div",{staticClass:"hero-body"},[a("div",{staticClass:"container"},[a("h1",{staticClass:"title"},[t._v("Twitter Popularity")]),t._v(" "),a("h2",{staticClass:"subtitle"},[t._v("Track the popularity of brands on twitter")])])])])}]}},47:function(t,e,a){a(114);var n=a(8)(a(68),a(128),null,null);t.exports=n.exports},48:function(t,e,a){"use strict";var n=a(25),s=a(132),i=a(125),r=a.n(i),o=a(124),l=a.n(o);n.a.use(s.a),e.a=new s.a({routes:[{path:"/",name:"Index",component:r.a},{path:"/detail/:id",name:"Detail",component:l.a}]})},49:function(t,e,a){a(117);var n=a(8)(a(69),a(131),null,null);t.exports=n.exports},67:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var n=a(25),s=a(49),i=a.n(s),r=a(48);n.a.config.productionTip=!1,new n.a({el:"#app",router:r.a,template:"<App/>",components:{App:i.a}})},68:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default={name:"PulseLoader",props:{loading:{type:Boolean,default:!0},color:{type:String,default:"#5dc596"},size:{type:String,default:"15px"},margin:{type:String,default:"2px"},radius:{type:String,default:"100%"}},data:function(){return{spinnerStyle:{backgroundColor:this.color,width:this.size,height:this.size,margin:this.margin,borderRadius:this.radius,display:"inline-block",animationName:"v-pulseStretchDelay",animationDuration:"0.75s",animationIterationCount:"infinite",animationTimingFunction:"cubic-bezier(.2,.68,.18,1.08)",animationFillMode:"both"},spinnerDelay1:{animationDelay:"0.12s"},spinnerDelay2:{animationDelay:"0.24s"},spinnerDelay3:{animationDelay:"0.36s"}}}}},69:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default={name:"app"}},70:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var n=a(33),s=a.n(n),i=a(32),r=a.n(i),o=a(26),l=a.n(o),c=a(126),u=a.n(c),d=a(47),p=a.n(d);e.default={components:{Plotly:u.a,PulseLoader:p.a},created:function(){var t=this,e=this.$route.params.id;this.loadTerm();var a=window.location.port;this.ws=new window.WebSocket("ws://127.0.0.1:"+a+"/ws/changes/"+e),this.ws.onopen=function(t){return console.log("opened websocket")},this.ws.onclose=function(t){return console.log("closed")},this.ws.onmessage=function(a){var n=JSON.parse(a.data);n.id===e&&(t.sum+=n.data.sentiment,t.length+=1,n.data.sentiment>.75?(t.chart[0].x.push(new Date(n.data.timestamp)),t.chart[0].y.push(1)):n.data.sentiment>.5?(t.chart[1].x.push(new Date(n.data.timestamp)),t.chart[1].y.push(.5)):n.data.sentiment>.25?(t.chart[2].x.push(new Date(n.data.timestamp)),t.chart[2].y.push(-.5)):n.data.sentiment<.25&&(t.chart[3].x.push(new Date(n.data.timestamp)),t.chart[3].y.push(-1)))}},beforeDestroy:function(){this.ws&&this.ws.close()},data:function(){return{loading:!1,selected:3600,sum:0,length:0,term:null,chart:null,ws:null}},watch:{selected:function(t){this.loadTerm(t)}},computed:{average:function(){return 0===this.length?0:this.sum/this.length}},methods:{chartData:function(t){var e={x:[],y:[],type:"bar",name:"love",marker:{color:"rgb(12, 134, 35)"}},a={x:[],y:[],type:"bar",name:"positive",marker:{color:"rgb(38, 173, 203)"}},n={x:[],y:[],type:"bar",name:"negative",marker:{color:"rgb(245, 139, 28)"}},s={x:[],y:[],type:"bar",name:"hate",marker:{color:"rgb(232, 20, 20)"}};return t.forEach(function(t){t.sentiment>.75?(e.x.push(new Date(t.time)),e.y.push(1)):t.sentiment>.5?(a.x.push(new Date(t.time)),a.y.push(.5)):t.sentiment>.25?(n.x.push(new Date(t.time)),n.y.push(-.5)):t.sentiment<.25&&(s.x.push(new Date(t.time)),s.y.push(-1))}),[e,a,n,s]},loadTerm:function(){var t=this,e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:3600;return r()(s.a.mark(function a(){var n,i,r;return s.a.wrap(function(a){for(;;)switch(a.prev=a.next){case 0:return n=t.$route.params.id,t.loading=!0,a.prev=2,a.next=5,l.a.get("api/terms/"+n,{params:{seconds:e}});case 5:i=a.sent,r=i.data,t.term=r.term,t.length=r.data.length,t.chart=t.chartData(r.data),t.sum=0===r.data.length?0:r.data.reduce(function(t,e){return{sentiment:t.sentiment+e.sentiment}}).sentiment,a.next=16;break;case 13:a.prev=13,a.t0=a.catch(2),console.log(a.t0);case 16:t.loading=!1;case 17:case"end":return a.stop()}},a,t,[[2,13]])}))()}}}},71:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var n=a(33),s=a.n(n),i=a(32),r=a.n(i),o=a(26),l=a.n(o),c=a(47),u=a.n(c);e.default={name:"Index",components:{PulseLoader:u.a},data:function(){return{loading:!1,term:"",terms:[]}},created:function(){var t=this;return r()(s.a.mark(function e(){var a,n;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return t.loading=!0,e.prev=1,e.next=4,l.a.get("api/terms");case 4:a=e.sent,n=a.data,t.terms=n,e.next=12;break;case 9:e.prev=9,e.t0=e.catch(1),console.log(e.t0);case 12:t.loading=!1;case 13:case"end":return e.stop()}},e,t,[[1,9]])}))()},methods:{goToDetails:function(t){this.$router.push({name:"Detail",params:{id:t}})},createTerm:function(){var t=this;return r()(s.a.mark(function e(){var a,n;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:if(""===t.term){e.next=13;break}return e.prev=1,e.next=4,l.a.post("api/terms",{term:t.term});case 4:a=e.sent,n=a.data,t.terms.push(n),t.term="",e.next=13;break;case 10:e.prev=10,e.t0=e.catch(1),console.log(e.t0);case 13:case"end":return e.stop()}},e,t,[[1,10]])}))()}}}},72:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var n=a(75),s=a.n(n),i=a(120),r=a.n(i);e.default={props:{data:{type:Array,required:!0,default:function(){return[]}},layout:{type:Object,default:function(){return{}}},config:{type:Object,default:function(){return{modeBarButtonsToRemove:["sendDataToCloud","hoverCompareCartesian"],displaylogo:!1}}},min:{type:Number,default:null},max:{type:Number,default:null},filename:{type:String,default:"newplot"}},mounted:function(){var t=this;r.a.plot(this.$el,this.data,this.layout,this.config).then(function(e){e.fn=t.filename,t.chart=e}),window.addEventListener("resize",this.handleResize)},beforeDestroy:function(){window.removeEventListener("resize",this.handleResize)},data:function(){return{chart:null}},computed:{finalLayout:function(){var t=s()({},this.layout,{shapes:[]});return null!==this.min&&t.shapes.push(this.horizontalLine(this.min,"#25A9E1")),null!==this.max&&t.shapes.push(this.horizontalLine(this.max,"#FFA100")),t}},methods:{horizontalLine:function(t){return{type:"line",xref:"paper",x0:0,y0:t,x1:1,y1:t,line:{color:arguments.length>1&&void 0!==arguments[1]?arguments[1]:"rgb(0, 0, 0)",width:4,dash:"dot"}}},handleResize:function(t){r.a.Plots.resize(this.chart)}},watch:{finalLayout:{handler:function(t){var e=this;this.$nextTick(function(){r.a.relayout(e.$refs.chart,t)})},immediate:!0}}}}},[67]);
//# sourceMappingURL=app.a4815cc1ad78f6d730db.js.map
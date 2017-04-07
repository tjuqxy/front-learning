var sudoData = {
    1: {
        3: 123456789,
        6: 12345678,
        9: 1234567,
    },
    4: {
        1: 123456,
        6: 12345,
        7: 1234,
    },
    8: {
        2: 123,
        3: 12,
        9: 1
    }
}

sudoData = {}

var app = new Vue({
    el: '#vue-demo-app1',
    data: {
        message: 'Hello Vue!'
    }
})

var app2 = new Vue({
    el: '#vue-demo-app2',
    data: {
        message: 'You loaded this page on ' + new Date()
    }
})

var app3 = new Vue({
    el: '#vue-demo-app3',
    data: {
        seen: true
    }
})

var app4 = new Vue({
    el: '#vue-demo-app4',
    data: {
        todos: [
            { text: 'Learn JavaScript' },
            { text: 'Learn Vue' },
            { text: 'Build something awesome' }
        ]
    }
})

var app5 = new Vue({
    el: '#vue-demo-app5',
    data: {
        message: 'Hello Vue!'
    },
    methods: {
        reverseMessage: function() {
            this.message = this.message.split('').reverse().join('')
        }
    }
})

var app6 = new Vue({
    el: '#vue-demo-app6',
    data: {
        message: 'Hello Vue'
    }
})

Vue.component('todo-item', {
  props: ['todo'],
  template: '<li>{{ todo.text }}</li>'
})

var app7 = new Vue({
    el: '#vue-demo-app7',
    data: {
        graceryList: [
            { text: 'a' },
            { text: 'cccc' },
            { text: 'b' }
        ]
    }
})

var app_vue = new Vue({
    el: '#vue-sudo',
    data: {
        numbers: sudoData
    },
    beforeCreate: function() {
        for(var row=1; row<10; row++) {
            if (sudoData[row] == null) {
                sudoData[row] = {}
            }
            for(var col=1; col<10; col++) {
                if (sudoData[row][col] == null) {
                    sudoData[row][col] = ""
                }
            }
        }
    },
    methods: {
        updateValue: function(row, col, numStr) {
            console.log("row: " + row)
            console.log("col: " + col)
            console.log("num: " + numStr)
            var valueStr = String(this.numbers[row][col])
            var index = valueStr.indexOf(numStr)
            if (index >= 0) {
                valueStr = valueStr.substring(0, index) + valueStr.substring(index+1)
            } else {
                valueStr += numStr
            }
            this.numbers[row][col] = valueStr.split("").sort().join("")
        },
        calculateSudo: function() {
            console.log("click calculateSudo")
            arr = {}
            for (var row=1; row<10; row++) {
                if (this.numbers[row] != null) {
                    arr[row] = {}
                    for (var col=1; col<10; col++) {
                        arr[row][col] = this.numbers[row][col]
                    }
                }
            }
            $.post(
                "/api/sudo/calculate",
                JSON.stringify({'sudo_num': arr}),
                function(data, status) {
                    console.log(data, status)
                    ret = JSON.parse(data)
                    console.log(ret)
                    for (var row=1; row<10; row++) {
                        for (var col=1; col<10; col++) {
                            app_vue.numbers[row][col] = ret[row-1][col-1]
                        }
                    }
                }
            )
        }
    },
    components: {
        'sudo-button': {
            props: ['value', 'row', 'col'],
            template: '<button @keyup="dealKeyup($event)" class="btn btn-default sudo-btn" :style="{color: computeColor, fontSize: computeFontSize, lineHeight: computeLineHeight}">{{ value }}</button>',
            methods: {
                dealKeyup: function(event) {
                    var numStr = String(event.keyCode - 48)
                    console.log('press num:' + numStr)
                    this.$emit("update", this.row, this.col, numStr)
                    return

                    var valueStr = String(this.value)
                    var index = valueStr.indexOf(numStr)
                    if (index >= 0) {
                        this.value = valueStr.substring(0, index) + valueStr.substring(index+1)
                    } else {
                        this.value += numStr
                    }
                }
            },
            computed: {
                valueLength: function() {
                    return String(this.value).length
                },
                computeColor: function() {
                    if (this.valueLength == 1) {
                        return 'dark'
                    }
                    return 'darkgray'
                },
                computeFontSize: function() {
                    switch (this.valueLength) {
                        case 1:
                            return '30px'
                        case 2:
                            return '25px'
                        case 3:
                            return '20px'
                        case 4:
                            return '22px'
                        case 5:
                            return '20px'
                        case 6:
                            return '18px'
                        case 7:
                        case 8:
                            return '15px'
                        case 9:
                            return '12px'
                        default:
                            return '20px'
                    }
                },
                computeLineHeight: function() {
                    switch (this.valueLength) {
                        case 4:
                        case 5:
                        case 6:
                            return '20px'
                        default:
                            return '20px'
                    }
                }
            }
        }
    }
})

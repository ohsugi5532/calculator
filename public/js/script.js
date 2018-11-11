const MAX_LENGTH = 29;

var app = new Vue ({
    el:"#app",　// DOM selector related to the instance of Vue.
    data:{
        result: "0",
        formula: "",
        error: "",
    },
    isStarted: false,               //
    isFractionalMode: false,        // Fractional Mode (ex. 0.001)
    isNumberInputed: false,         // Any number key inputed
    isOperatorInputed: false,       // Any operator (ex. +) inputed
    isPointInputed: false,          // Decimal point inputed
    isLeftBracketInputed: false,    // Left bracket "( inputed
    isRightBracketInputed: false,   // Roght bracket ")" inputed
    hasSent: false,
    bracketCount: 0,
    sendFormula: "",
    methods:{
        enter: function() {
            // ajax process uisng axios
            // option definition in requesting.
            config = {
                headers:{
                    'X-Requested-With': 'XMLHttpRequest',
                    'Content-Type':'application/x-www-form-urlencoded'
                },
                withCredentials:false,
            }

            // create POST parameter
            var params = new URLSearchParams();
            params.append('formula', this.sendFormula);

            // ajax request
            axios.post("/", params, config)
            .then(function(res) {
                app.result = res.data
                console.log(res)
            })
            .catch(error => {
                console.log(error)
            })
            .finally(() => {
                this.isFractionalMode = false;
                this.isNumberInputed = false;
                this.isOperatorInputed = false;
                this.isPointInputed = false;
                this.isLeftBracketInputed = false;
                this.isRightBracketInputed = false;

                this.sendFormula = ""
                this.hasSent = true;
            })
        },
        setNumber: function(number) {
            if (this.hasSent == true) {
                this.formula = "";
                this.result = "";
                this.hasSent = false;
            }

            if (this.formula.length >= MAX_LENGTH) {
                this.result = "TOO LONG EXPRESSION"
                return;
            }

            if (this.sendFormula == undefined) {
                this.sendFormula = "";
            }

            if (this.isRightBracketInputed != true) {
                this.isNumberInputed = true;
                this.isPointInputed = false;
                this.isOperatorInputed = false;
                this.isLeftBracketInputed = false;

                this.formula += number + "";
                this.sendFormula += number + "";
            }
        },
        setOperator: function(operator) {
            if (this.hasSent == true) {
                this.formula = "";
                this.result = "";
                this.hasSent = false;
            }

            if (this.formula.length >= MAX_LENGTH) {
                this.result = "TOO LONG EXPRESSION"
                return;
            }

            if (this.isNumberInputed == true || this.isRightBracketInputed == true) {
                this.isOperatorInputed = true;
                this.isNumberInputed = false;
                this.isRightBracketInputed = false;

                // Change the fractional mode
                this.isFractionalMode = false;

                this.formula += operator;
                this.sendFormula += " " + operator + " ";
            }
        },
        setPoint: function() {
            if (this.hasSent == true) {
                this.formula = "";
                this.result = "";
                this.hasSent = false;
            }

            if (this.formula.length >= MAX_LENGTH) {
                this.result = "TOO LONG EXPRESSION"
                return;
            }

            if (this.isNumberInputed == true) {
                if (this.isFractionalMode != true) {
                    this.isPointInputed = true;

                    // Change the fractional mode
                    this.isFractionalMode = true;

                    this.formula += ".";
                    this.sendFormula += "."
                }
            }
        },
        setLeftBracket: function() {
            if (this.hasSent == true) {
                this.formula = "";
                this.result = "";
                this.hasSent = false;
            }

            if (this.formula.length >= MAX_LENGTH) {
                this.result = "TOO LONG EXPRESSION"
                return;
            }

            if (this.sendFormula == undefined) {
                this.sendFormula = "";
            }

            if (this.formula == "" || this.isOperatorInputed == true || this.isLeftBracketInputed == true) {
                this.isLeftBracketInputed = true;
                this.isOperatorInputed = false;

                this.formula += "(";
                this.sendFormula += " ( ";
            }
        },
        setRightBracket: function() {
            if (this.hasSent == true) {
                this.formula = "";
                this.result = "";
                this.hasSent = false;
            }

            if (this.formula.length >= MAX_LENGTH) {
                this.result = "TOO LONG EXPRESSION"
                return;
            }

            if (this.isNumberInputed == true || this.isRightBracketInputed == true) {
                this.isRightBracketInputed = true;
                this.isNumberInputed = false;

                this.formula += ")";
                this.sendFormula += " ) ";
            }
        },
        clear: function() {
            this.formula = "";
            this.result = "0";
            this.error = "";
            this.sendFormula = "";
            this.isAnyInputed = false;
            this.isFractionalMode = false;
            this.isPointInputed = false;
            this.isLeftBracketInputed = false;
            this.isRightBracketInputed = false;
            this.isOperatorInputed = false;
        }
    }
})
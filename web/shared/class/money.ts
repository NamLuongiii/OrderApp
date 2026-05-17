import BigNumber from "bignumber.js";

export default class Money {

    constructor(public amount: string | number) {
        this.amount = amount;
    }

    plus(m: string | number): Money {
        return new Money(new BigNumber(this.amount).plus(m).toString())
    }

    multiply(m: string | number): Money {
        return new Money(new BigNumber(this.amount).multipliedBy(m).toString())
    }

    format(): string {
        return new BigNumber(this.amount).toFormat({ suffix: 'đ' })
    }

    toString(): string {
        return this.amount.toString()
    }
}
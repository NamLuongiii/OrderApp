export default class Time {
    private unixTime: number;

    constructor(unixTime: number) {
        this.unixTime = unixTime;
    }

    format() {
        return new Date(this.unixTime).toLocaleString()
    }
}
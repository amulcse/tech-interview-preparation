class Heap {
    items = [];

    push(item) {
        this.items.push(item)
        this.bubbleUp(this.items.length - 1)
    }

    pop() {

        if (this.items.length == 0) {
            return -1
        }

        const item = this.items[0]
        const last = this.items.pop();

        if (this.items.length > 0) {
            this.items[0] = last;
            this.bubbleDown(0);
        }

        return item

    }

    peek() {

        if (this.items.length > 0) {
            return this.items[0]
        }

        return -1

    }

    bubbleUp(index) {

        if (index <= 0) {
            return
        }

        let parent = Math.floor((index - 1) / 2)

        if (this.items[parent] < this.items[index]) {
            let temp = this.items[index]
            this.items[index] = this.items[parent]
            this.items[parent] = temp

            this.bubbleUp(parent)
        }

    }

    bubbleDown(index) {

        if (index == this.items.length) {
            return
        }

        let c1 = 2 * index + 1
        let c2 = 2 * index + 2

        if (this.items[c1] > this.items[c2]) {

            if (this.items[c1] > this.items[index]) {
                let temp = this.items[index]
                this.items[index] = this.items[c1]
                this.items[c1] = temp
                this.bubbleDown(c1)
            }

        } else {

            if (this.items[c2] > this.items[index]) {
                let temp = this.items[index]
                this.items[index] = this.items[c2]
                this.items[c2] = temp
                this.bubbleDown(c2)
            }

        }


    }
}

let h = new Heap()

h.push(10)
h.push(16)
h.push(18)
h.push(16)
h.push(10)
h.push(23)


console.log(h.pop())
console.log(h.pop())
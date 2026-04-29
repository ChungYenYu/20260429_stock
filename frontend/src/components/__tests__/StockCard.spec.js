import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import StockCard from '../StockCard.vue'

describe('StockCard.vue', () => {
  it('renders props correctly', () => {
    const stock = {
      code: '2330',
      name: 'TSMC',
      price: 1000
    }
    const wrapper = mount(StockCard, {
      props: { stock }
    })
    expect(wrapper.text()).toContain('2330')
    expect(wrapper.text()).toContain('TSMC')
    expect(wrapper.text()).toContain('1,000')
  })
})

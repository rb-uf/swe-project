import { SubmitButtonComponent } from './submit-button.component'

describe('SubmitButtonComponent', () => {
  it('mounts', () => {
    cy.mount(SubmitButtonComponent)
  })
  it('should ', () => {
    cy.mount(SubmitButtonComponent)
    cy.get('[data-cy=counter]').should('have.text', '0')
  })
})
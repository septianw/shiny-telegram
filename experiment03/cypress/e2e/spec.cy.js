describe('template spec', () => {
  it('passes', () => {
    cy.visit('http://localhost:5174')
    cy.contains('count is').should('have.text', 'count is 0')
    cy.get("#app > div.card > button").click()
    cy.contains('count is').should('have.text', 'count is 1')
  })
})